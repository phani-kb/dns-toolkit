package processors

import (
	"net/url"
	"strings"
)

const (
	AdguardExceptionPrefix = "@@"
	AdguardBlockPrefix     = "||"
	AdguardBlockSuffix     = "^"
	AdguardBlockSuffixFull = "$"
)

// ParseToAdguardEntry parses a raw URL string and returns an AdGuard folder-style entry
func ParseToAdguardEntry(raw string) (string, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", false
	}
	parsed, err := url.Parse(raw)
	if err != nil {
		return "", false
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", false
	}
	if parsed.Host == "" {
		return "", false
	}
	rest := parsed.Host + parsed.RequestURI()
	entry := AdguardBlockPrefix + rest + "$"
	return entry, true
}

func EqualUnorderedAdg(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
