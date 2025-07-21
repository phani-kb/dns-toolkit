package processors

import (
	"testing"

	"github.com/phani-kb/multilog"
)

func TestIpv4UrlProcessor_Process(t *testing.T) {
	tests := []struct {
		name    string
		content string
		valid   []string
		invalid []string
	}{
		{
			name:    "valid ipv4 urls",
			content: "192.168.1.100/auth/login\n192.168.1.200/auth/login",
			valid:   []string{"192.168.1.100", "192.168.1.200"},
			invalid: []string{},
		},
		{
			name:    "invalid entries",
			content: "not-an-ip/auth/login\n256.256.256.256/auth/login",
			valid:   []string{},
			invalid: []string{"not-an-ip/auth/login", "256.256.256.256/auth/login"},
		},
		{
			name:    "mixed valid and invalid",
			content: "192.168.1.100/auth/login\ninvalid-url\n192.168.1.200/auth/login",
			valid:   []string{"192.168.1.100", "192.168.1.200"},
			invalid: []string{"invalid-url"},
		},
		{
			name:    "domain url entry",
			content: "example.com/auth/login",
			valid:   []string{},
			invalid: []string{"example.com/auth/login"},
		},
	}

	processor := NewIpv4UrlProcessor("ipv4_url", "test")
	logger := &multilog.Logger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, invalid := processor.Process(logger, tt.content)
			if len(valid) != len(tt.valid) {
				t.Errorf("Expected %d valid entries, got %d", len(tt.valid), len(valid))
			}
			for i, v := range tt.valid {
				if i >= len(valid) || valid[i] != v {
					t.Errorf("Expected valid[%d] to be %s, got %s", i, v, valid[i])
				}
			}
			if len(invalid) != len(tt.invalid) {
				t.Errorf("Expected %d invalid entries, got %d", len(tt.invalid), len(invalid))
			}
			for i, v := range tt.invalid {
				if i >= len(invalid) || invalid[i] != v {
					t.Errorf("Expected invalid[%d] to be %s, got %s", i, v, invalid[i])
				}
			}
		})
	}
}
