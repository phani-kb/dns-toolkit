package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewDomainCommentProcessor(t *testing.T) {
	t.Parallel()

	sourceType := "domain_comment"
	listType := "blocklist"

	processor := processors.NewDomainCommentProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestDomainCommentProcessor_Process(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCommentProcessor("domain_comment", "blocklist")

	tests := []struct {
		name            string
		content         string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "empty content",
			content:         "",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "single IP with domain comment",
			content:         "1.0.0.1             # 1dot1dot1dot1.cloudflare-dns.com",
			expectedValid:   []string{"1dot1dot1dot1.cloudflare-dns.com"},
			expectedInvalid: nil,
		},
		{
			name:            "IP with domain comment and trailing dot",
			content:         "1.1.1.1             # example.com.",
			expectedValid:   []string{"example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "multiple IP entries with domains",
			content:         "1.0.0.1             # cloudflare-dns.com\n1.1.1.1             # example.org\n8.8.8.8             # dns.google",
			expectedValid:   []string{"cloudflare-dns.com", "example.org", "dns.google"},
			expectedInvalid: nil,
		},
		{
			name:            "mixed valid and invalid lines",
			content:         "1.0.0.1             # cloudflare-dns.com\ninvalid line\n1.1.1.1             # example.org\njust text",
			expectedValid:   []string{"cloudflare-dns.com", "example.org"},
			expectedInvalid: []string{"invalid line", "just text"},
		},
		{
			name:            "IP with comment but no valid domain",
			content:         "1.0.0.1             # not a valid domain name",
			expectedValid:   nil,
			expectedInvalid: []string{"1.0.0.1             # not a valid domain name"},
		},
		{
			name:            "comment lines should be ignored",
			content:         "# This is a header comment\n! Another comment\n// Yet another comment",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "IP with invalid domain in comment",
			content:         "1.0.0.1             # invalid-domain-123\n8.8.8.8             # .invalid\n9.9.9.9             # valid.example.com",
			expectedValid:   []string{"valid.example.com"},
			expectedInvalid: []string{"1.0.0.1             # invalid-domain-123", "8.8.8.8             # .invalid"},
		},
		{
			name:    "real sample data",
			content: "1.0.0.1             # 1dot1dot1dot1.cloudflare-dns.com\n1.0.0.2             # security.cloudflare-dns.com\n1.1.1.1             # 1dot1dot1dot1.cloudflare-dns.com\n3.33.139.32         # resolver.unstoppable.io",
			expectedValid: []string{
				"1dot1dot1dot1.cloudflare-dns.com",
				"security.cloudflare-dns.com",
				"1dot1dot1dot1.cloudflare-dns.com",
				"resolver.unstoppable.io",
			},
			expectedInvalid: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, invalid := processor.Process(logger, tt.content)
			assert.Equal(t, tt.expectedValid, valid, "Valid entries should match expected")
			assert.Equal(t, tt.expectedInvalid, invalid, "Invalid entries should match expected")
		})
	}
}
