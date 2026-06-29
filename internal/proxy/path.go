package proxy

import "strings"

// NormalizeAPIPath ensures routing uses the full gateway prefix. Chi sub-routers
// may present paths with the /api/v1 mount stripped (e.g. /payments/ledger/...).
func NormalizeAPIPath(path string) string {
	if strings.HasPrefix(path, "/api/v1/") || path == "/api/v1" {
		return path
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return "/api/v1" + path
}
