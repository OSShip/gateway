package proxy

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"github.com/OSShip/utils/observability"
	"github.com/redis/go-redis/v9"
)

func CacheKeyForRequest(r *http.Request) string {
	path := r.URL.Path
	q := r.URL.Query()
	if path == "/api/v1/listings" && q.Get("status") == "active" && q.Get("oss_project") == "" {
		return "listings:active"
	}
	if path == "/api/v1/public/payout-summary" {
		return "public:ledger:summary"
	}
	if strings.HasPrefix(path, "/api/v1/listings/") && len(strings.Split(strings.TrimPrefix(path, "/api/v1/listings/"), "/")) == 1 {
		return "listings:id:" + strings.TrimPrefix(path, "/api/v1/listings/")
	}
	if path == "/api/v1/public/listings" && q.Get("status") == "active" && q.Get("oss_project") == "" {
		return "listings:active"
	}
	h := sha256.Sum256([]byte(r.URL.String()))
	return "cache:" + hex.EncodeToString(h[:8])
}

func CacheLabel(key string) string {
	if strings.HasPrefix(key, "listings:") || strings.HasPrefix(key, "public:") {
		return key
	}
	return "other"
}

func IsCacheableGET(path string) bool {
	return strings.HasPrefix(path, "/api/v1/listings") || strings.HasPrefix(path, "/api/v1/public")
}

func CacheTTL(path string) time.Duration {
	if strings.Contains(path, "payout-summary") {
		return 300 * time.Second
	}
	return 60 * time.Second
}

func InvalidateListingCache(ctx context.Context, rdb *redis.Client) {
	keys := []string{"listings:active", "public:ledger:summary"}
	_ = rdb.Del(ctx, keys...).Err()
	iter := rdb.Scan(ctx, 0, "listings:id:*", 100).Iterator()
	for iter.Next(ctx) {
		_ = rdb.Del(ctx, iter.Val()).Err()
	}
	iter = rdb.Scan(ctx, 0, "cache:*", 100).Iterator()
	for iter.Next(ctx) {
		_ = rdb.Del(ctx, iter.Val()).Err()
	}
}

func RecordCacheHit(label string) {
	observability.CacheHits.WithLabelValues(label).Inc()
}

func RecordCacheMiss(label string) {
	observability.CacheMisses.WithLabelValues(label).Inc()
}
