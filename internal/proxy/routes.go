package proxy

import "strings"

func ResolveBackend(path string, backends map[string]string) (string, string) {
	routes := []struct {
		prefix  string
		service string
		strip   string
	}{
		{"/api/v1/auth", "auth", "/api/v1/auth"},
		{"/api/v1/listings", "listings", "/api/v1/listings"},
		{"/api/v1/users", "users", "/api/v1/users"},
		{"/api/v1/sessions", "sessions", "/api/v1/sessions"},
		{"/api/v1/mentors", "mentors", "/api/v1/mentors"},
		{"/api/v1/payments", "payments", "/api/v1/payments"},
		{"/api/v1/public/listings", "listings", "/api/v1/public/listings"},
		{"/api/v1/public/payout-summary", "payments", "/api/v1/public/payout-summary"},
	}
	for _, rt := range routes {
		if strings.HasPrefix(path, rt.prefix) {
			return backends[rt.service], rt.strip
		}
	}
	return "", ""
}

func RewritePath(stripped, stripPrefix string) string {
	if stripPrefix == "/api/v1/public/payout-summary" {
		return "/payout-summary"
	}
	if stripPrefix == "/api/v1/public/listings" {
		if stripped == "" {
			return "/"
		}
		return stripped
	}
	return stripped
}
