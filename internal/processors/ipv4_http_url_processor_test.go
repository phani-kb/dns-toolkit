package processors

import (
	"testing"

	"github.com/phani-kb/multilog"
)

func TestIpv4HttpUrlProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := NewIpv4HttpUrlProcessor("ipv4_http_url", "test")

	tests := []struct {
		name    string
		input   string
		want    []string
		wantInv []string
	}{
		{
			name:    "valid http url",
			input:   "http://192.168.1.1/path",
			want:    []string{"192.168.1.1"},
			wantInv: []string{},
		},
		{
			name:    "valid https url",
			input:   "https://8.8.8.8:8080/test",
			want:    []string{"8.8.8.8"},
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
			input:   "# comment\n\nhttp://10.0.0.1",
			want:    []string{"10.0.0.1"},
			wantInv: []string{},
		},
		{
			name:    "invalid ipv4 in url",
			input:   "http://999.999.999.999/path",
			want:    []string{},
			wantInv: []string{"http://999.999.999.999/path"},
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
