// @title           OSShip API
// @version         1.0
// @description     OSShip open-source mentorship platform API exposed via the gateway.
// @host            localhost
// @BasePath        /api/v1
//
// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 JWT token. Example: Bearer {token}
package main

import (
	"net/http"
	"os"
	"time"

	"github.com/OSShip/gateway/internal/config"
	"github.com/OSShip/gateway/internal/proxy"
	"github.com/OSShip/gateway/internal/ratelimit"
	"github.com/OSShip/utils/observability"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "github.com/OSShip/gateway/docs"
	_ "github.com/OSShip/gateway/internal/apidoc"
)

func main() {
	cfg := config.Load()
	observability.InitSentry("gateway")
	defer observability.FlushSentry(2 * time.Second)
	logger := observability.InitLogger("gateway")

	opt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		logger.Error("redis parse failed", "err", err)
		os.Exit(1)
	}
	rdb := redis.NewClient(opt)
	logger.Info("redis connected", "url", cfg.RedisURL)
	limiter := redis_rate.NewLimiter(rdb)
	routeLimits := ratelimit.ApplyOverrides(observability.DefaultRouteLimits())

	proxyHandler := &proxy.Handler{
		Backends:    cfg.Backends,
		Redis:       rdb,
		Limiter:     limiter,
		JWTSecret:   cfg.JWTSecret,
		RouteLimits: routeLimits,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(observability.SentryHTTPMiddleware)
	r.Use(observability.SentryRecoverMiddleware("gateway"))
	r.Use(observability.SentryErrorMiddleware("gateway"))
	r.Use(observability.RequestLogMiddleware("gateway"))
	r.Use(observability.PrometheusMiddleware("gateway"))

	r.Get("/health", observability.HealthHandler("gateway"))
	r.Get("/metrics", observability.MetricsHandler().ServeHTTP)
	r.Get("/api/v1/health", observability.HealthHandler("gateway"))

	r.Get("/api/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/docs/index.html", http.StatusMovedPermanently)
	})
	r.Get("/api/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/api/docs/doc.json"),
	))

	// Single handler for all /api/v1 traffic (avoid chi sub-router path stripping issues).
	r.Handle("/api/v1", proxyHandler)
	r.Handle("/api/v1/*", proxyHandler)

	logger.Info("gateway listening", "port", cfg.Port, "backends", len(cfg.Backends), "docs", "/api/docs/")
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		logger.Error("server failed", "err", err)
		os.Exit(1)
	}
}
