package proxy

import "testing"

func TestNormalizeAPIPath(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"/api/v1/payments/ledger/abc", "/api/v1/payments/ledger/abc"},
		{"/payments/ledger/abc", "/api/v1/payments/ledger/abc"},
		{"/api/v1", "/api/v1"},
		{"/auth/login", "/api/v1/auth/login"},
	}
	for _, tc := range tests {
		if got := NormalizeAPIPath(tc.in); got != tc.want {
			t.Errorf("NormalizeAPIPath(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestResolveBackendPaymentsLedger(t *testing.T) {
	backends := map[string]string{"payments": "http://payments:8087"}
	target, strip := ResolveBackend("/api/v1/payments/ledger/listing-id", backends)
	if target != "http://payments:8087" || strip != "/api/v1/payments" {
		t.Fatalf("got target=%q strip=%q", target, strip)
	}
	stripped := NormalizeAPIPath("/api/v1/payments/ledger/listing-id")
	stripped = stripped[len(strip):]
	if stripped != "/ledger/listing-id" {
		t.Fatalf("stripped path = %q", stripped)
	}
	if got := RewritePath(stripped, strip); got != "/ledger/listing-id" {
		t.Fatalf("rewrite = %q", got)
	}
}
