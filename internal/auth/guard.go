package auth

import (
	"net/http"
	"strings"
)

var protectedRoutes = []struct {
	method string
	prefix string
}{
	{http.MethodGet, "/api/v1/auth/me"},
	{http.MethodPatch, "/api/v1/users/me"},
	{http.MethodPost, "/api/v1/users/me/"},
}

func RequiresAuth(path, method string) bool {
	for _, pr := range protectedRoutes {
		if method != pr.method {
			continue
		}
		if strings.HasSuffix(pr.prefix, "/") {
			if strings.HasPrefix(path, pr.prefix) {
				return true
			}
		} else if path == pr.prefix || strings.HasPrefix(path, pr.prefix+"/") {
			return true
		}
	}
	return false
}
