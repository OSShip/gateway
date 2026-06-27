package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/OSShip/utils/jwtutil"
	"github.com/redis/go-redis/v9"
)

type cachedClaims struct {
	UserID         string `json:"sub"`
	Role           string `json:"role"`
	GithubUsername string `json:"github_username,omitempty"`
}

func ResolveClaims(ctx context.Context, r *http.Request, rdb *redis.Client, jwtSecret string) (*jwtutil.Claims, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil, fmt.Errorf("missing authorization")
	}
	tokenStr := strings.TrimPrefix(auth, "Bearer ")
	if tokenStr == auth {
		return nil, fmt.Errorf("invalid authorization header")
	}

	hash := sha256.Sum256([]byte(tokenStr))
	cacheKey := "auth:session:" + hex.EncodeToString(hash[:16])

	if cached, err := rdb.Get(ctx, cacheKey).Bytes(); err == nil {
		var cc cachedClaims
		if json.Unmarshal(cached, &cc) == nil {
			return &jwtutil.Claims{UserID: cc.UserID, Role: cc.Role, GithubUsername: cc.GithubUsername}, nil
		}
	}

	claims, err := jwtutil.ValidateToken(jwtSecret, tokenStr)
	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt != nil {
		ttl := time.Until(claims.ExpiresAt.Time)
		if ttl > 0 {
			payload, _ := json.Marshal(cachedClaims{
				UserID:         claims.UserID,
				Role:           claims.Role,
				GithubUsername: claims.GithubUsername,
			})
			_ = rdb.Set(ctx, cacheKey, payload, ttl).Err()
		}
	}
	return claims, nil
}

func ClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return strings.TrimSpace(strings.Split(xff, ",")[0])
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}
