package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port      string
	JWTSecret string
	RedisURL  string
	Backends  map[string]string
}

func Load() Config {
	return Config{
		Port:      env("PORT", "8080"),
		JWTSecret: env("JWT_SECRET", "dev-secret"),
		RedisURL:  env("REDIS_URL", "redis://redis:6379"),
		Backends: map[string]string{
			"auth":     env("AUTH_SERVICE_URL", "http://auth:8081"),
			"listings": env("LISTINGS_SERVICE_URL", "http://listings:8082"),
			"users":    env("USERS_SERVICE_URL", "http://users:8083"),
			"sessions": env("SESSIONS_SERVICE_URL", "http://sessions:8084"),
			"mentors":  env("MENTORS_SERVICE_URL", "http://mentors:8085"),
			"payments": env("PAYMENTS_SERVICE_URL", "http://payments:8087"),
		},
	}
}

func EnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
