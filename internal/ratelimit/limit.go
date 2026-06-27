package ratelimit

import (
	"context"
	"net/http"

	"github.com/go-redis/redis_rate/v10"
	"github.com/OSShip/utils/observability"
	gatewayauth "github.com/OSShip/gateway/internal/auth"
	"github.com/OSShip/gateway/internal/config"
	"github.com/redis/go-redis/v9"
)

func ApplyOverrides(rules []observability.RouteLimitRule) []observability.RouteLimitRule {
	out := make([]observability.RouteLimitRule, len(rules))
	copy(out, rules)
	if v := config.EnvInt("RATE_LIMIT_AUTH_LOGIN", 0); v > 0 {
		limit := redis_rate.PerMinute(v)
		for i := range out {
			if out[i].Limit.Group == "auth_login" {
				out[i].Limit.Limit = limit
			}
		}
	}
	if v := config.EnvInt("RATE_LIMIT_PAYMENTS_CHECKOUT", 0); v > 0 {
		limit := redis_rate.PerMinute(v)
		for i := range out {
			if out[i].Limit.Group == "payments_checkout" {
				out[i].Limit.Limit = limit
			}
		}
	}
	return out
}

func Check(ctx context.Context, r *http.Request, limiter *redis_rate.Limiter, jwtSecret string, rdb *redis.Client, rl observability.RouteLimit) (bool, int, error) {
	identifier := gatewayauth.ClientIP(r)
	if rl.ByUser {
		if claims, err := gatewayauth.ResolveClaims(ctx, r, rdb, jwtSecret); err == nil {
			identifier = claims.UserID
		}
	}
	key := observability.RateLimitKey(rl.Group, identifier)
	return observability.AllowRequest(ctx, limiter, key, rl.Limit)
}
