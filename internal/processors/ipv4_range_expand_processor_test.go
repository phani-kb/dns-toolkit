package processors

import (
	"testing"

	"github.com/phani-kb/multilog"
)

func TestIpv4RangeProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := NewIpv4RangeProcessor("ipv4_range_expand", "test")

	tests := []struct {
		name        string
		input       string
		validIP     string
		wantInvalid bool
	}{
		{"Valid range 192.168.1.0-192.168.1.255", "192.168.1.0-192.168.1.255", "192.168.1.100", false},
		{"Invalid range format", "192.168.1.0/24", "", true},
		{"Invalid IPs", "999.999.999.0-999.999.999.255", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, invalid := processor.Process(logger, tt.input)
			if tt.wantInvalid {
				if len(valid) != 0 {
					t.Errorf("Expected no valid entries, got: %v", valid)
				}
				if len(invalid) == 0 {
					t.Errorf("Expected invalid entries, got none")
				}
			} else {
				if len(invalid) != 0 {
					t.Errorf("Expected no invalid entries, got: %v", invalid)
				}
				found := false
				for _, ip := range valid {
					if ip == tt.validIP {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected IP %s in valid entries", tt.validIP)
				}
			}
		})
	}
}
