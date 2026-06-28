package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis_rate/v10"
	"github.com/OSShip/gateway/internal/config"
	"github.com/OSShip/gateway/internal/proxy"
	"github.com/OSShip/gateway/internal/ratelimit"
	"github.com/OSShip/utils/observability"
	"github.com/redis/go-redis/v9"
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

	r.Route("/api/v1", func(api chi.Router) {
		api.HandleFunc("/*", proxyHandler.ServeHTTP)
	})

	logger.Info("gateway listening", "port", cfg.Port, "backends", len(cfg.Backends))
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		logger.Error("server failed", "err", err)
		os.Exit(1)
	}
}
