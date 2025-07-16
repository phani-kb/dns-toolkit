package processors

import (
	"testing"

	"github.com/phani-kb/multilog"
)

func TestDomainHttpUrlProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := NewDomainHttpUrlProcessor("domain_http_url", "test")

	tests := []struct {
		name    string
		input   string
		want    []string
		wantInv []string
	}{
		{
			name:    "valid http url",
			input:   "http://example.com/path",
			want:    []string{"example.com"},
			wantInv: []string{},
		},
		{
			name:    "valid https url",
			input:   "https://sub.domain.org",
			want:    []string{"sub.domain.org"},
			wantInv: []string{},
		},
		{
			name:    "invalid line",
			input:   "not_a_url",
			want:    []string{},
			wantInv: []string{"not_a_url"},
		},
		{
			name:    "comment and empty line",
			input:   "# comment\n\nhttp://valid.com",
			want:    []string{"valid.com"},
			wantInv: []string{},
		},
		{
			name:    "duplicate domains",
			input:   "http://dup.com\nhttps://dup.com",
			want:    []string{"dup.com"},
			wantInv: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotInv := processor.Process(logger, tt.input)
			if !equalUnordered(got, tt.want) {
				t.Errorf("valid entries: got %v, want %v", got, tt.want)
			}
			if !equalUnordered(gotInv, tt.wantInv) {
				t.Errorf("invalid entries: got %v, want %v", gotInv, tt.wantInv)
			}
		})
	}
}

func equalUnordered(a, b []string) bool {
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
