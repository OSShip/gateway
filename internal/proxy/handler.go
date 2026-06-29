package proxy

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis_rate/v10"
	"github.com/OSShip/gateway/internal/auth"
	"github.com/OSShip/gateway/internal/ratelimit"
	"github.com/OSShip/utils/observability"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	Backends    map[string]string
	Redis       *redis.Client
	Limiter     *redis_rate.Limiter
	JWTSecret   string
	RouteLimits []observability.RouteLimitRule
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := NormalizeAPIPath(r.URL.Path)
	log := observability.DefaultLog

	if auth.RequiresAuth(path, r.Method) {
		if _, err := auth.ResolveClaims(r.Context(), r, h.Redis, h.JWTSecret); err != nil {
			log.WarnContext(r.Context(), "proxy auth rejected", "path", path, "method", r.Method, "err", err)
			observability.HTTPRequestsTotal.WithLabelValues("gateway", r.Method, path, "401").Inc()
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}
	}

	rl := observability.MatchRateLimit(r, h.RouteLimits)
	if allowed, retryAfter, err := ratelimit.Check(r.Context(), r, h.Limiter, h.JWTSecret, h.Redis, rl); err == nil && !allowed {
		log.WarnContext(r.Context(), "rate limit exceeded", "path", path, "group", rl.Group, "retry_after", retryAfter)
		observability.RateLimitExceeded.WithLabelValues("gateway", rl.Group).Inc()
		w.Header().Set("Retry-After", strconv.Itoa(retryAfter))
		http.Error(w, `{"error":"rate limit exceeded"}`, http.StatusTooManyRequests)
		return
	}

	target, stripPrefix := ResolveBackend(path, h.Backends)
	if target == "" {
		log.WarnContext(r.Context(), "no backend for path", "path", path)
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}

	log.DebugContext(r.Context(), "proxy routing", "path", path, "target", target, "strip_prefix", stripPrefix)

	if r.Method == http.MethodGet && IsCacheableGET(path) {
		cacheKey := CacheKeyForRequest(r)
		label := CacheLabel(cacheKey)
		if cached, err := h.Redis.Get(r.Context(), cacheKey).Bytes(); err == nil {
			log.DebugContext(r.Context(), "cache hit", "path", path, "label", label)
			RecordCacheHit(label)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Cache", "HIT")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(cached)
			return
		}
		RecordCacheMiss(label)
		log.DebugContext(r.Context(), "cache miss", "path", path, "label", label)
		rec := &responseRecorder{ResponseWriter: w, status: 200, body: &bytes.Buffer{}}
		h.forward(rec, r, target, stripPrefix)
		if rec.status == http.StatusOK && rec.body.Len() > 0 {
			if err := h.Redis.Set(r.Context(), cacheKey, rec.body.Bytes(), CacheTTL(path)).Err(); err != nil {
				log.WarnContext(r.Context(), "cache set failed", "path", path, "err", err)
			} else {
				log.DebugContext(r.Context(), "response cached", "path", path, "bytes", rec.body.Len())
			}
		}
		return
	}

	if r.Method != http.MethodGet && strings.HasPrefix(path, "/api/v1/listings") {
		defer InvalidateListingCache(r.Context(), h.Redis)
		log.DebugContext(r.Context(), "listing mutation, cache invalidated", "path", path)
	}

	h.forward(w, r, target, stripPrefix)
}

type responseRecorder struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func (r *responseRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (h *Handler) forward(w http.ResponseWriter, r *http.Request, targetURL, stripPrefix string) {
	target, err := url.Parse(targetURL)
	if err != nil {
		slog.ErrorContext(r.Context(), "bad backend URL", "target", targetURL, "err", err)
		http.Error(w, "bad gateway", http.StatusBadGateway)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		fullPath := NormalizeAPIPath(req.URL.Path)
		stripped := strings.TrimPrefix(fullPath, stripPrefix)
		req.URL.Path = RewritePath(stripped, stripPrefix)
		if req.URL.Path == "" {
			req.URL.Path = "/"
		}
		req.Host = target.Host

		if claims, err := auth.ResolveClaims(req.Context(), r, h.Redis, h.JWTSecret); err == nil {
			req.Header.Set("X-User-Id", claims.UserID)
			req.Header.Set("X-User-Role", claims.Role)
			if claims.GithubUsername != "" {
				req.Header.Set("X-Github-Username", claims.GithubUsername)
			}
		}
		if reqID := middleware.GetReqID(r.Context()); reqID != "" {
			req.Header.Set("X-Request-Id", reqID)
		}
	}

	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, proxyErr error) {
		slog.ErrorContext(req.Context(), "upstream unavailable",
			"target", targetURL,
			"path", req.URL.Path,
			"err", proxyErr,
		)
		observability.CaptureError(proxyErr, map[string]string{
			"service": "gateway",
			"target":  targetURL,
			"path":    req.URL.Path,
		})
		http.Error(rw, `{"error":"service unavailable"}`, http.StatusBadGateway)
	}
	proxy.ServeHTTP(w, r)
}
