package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name string
		input string
		want string
	}{
		{name: "uppercase HTTPS scheme", input: "HTTPS://example.com", want: "https://example.com"},
		{name: "mixed-case HTTP scheme", input: "HtTp://example.com", want: "http://example.com"},
		{name: "missing scheme defaults to HTTPS", input: "example.com", want: "https://example.com"},
		{name: "surrounding whitespace is trimmed", input: "  example.com  ", want: "https://example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeURL(tt.input); got != tt.want {
				t.Fatalf("normalizeURL(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
