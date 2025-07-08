package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewHostnameProcessor(t *testing.T) {
	t.Parallel()

	sourceType := "hostname"
	listType := "blocklist"

	processor := processors.NewHostnameProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestHostnameProcessor_Process(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewHostnameProcessor("hostname", "blocklist")

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
			name:            "single valid entry",
			content:         "127.0.0.1 example.com",
			expectedValid:   []string{"example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "multiple valid entries",
			content:         "127.0.0.1 example.com\n0.0.0.0 malicious.com\n127.0.0.1 ads.tracker.com",
			expectedValid:   []string{"example.com", "malicious.com", "ads.tracker.com"},
			expectedInvalid: nil,
		},
		{
			name:            "entries with IPv6",
			content:         "::1 localhost\n127.0.0.1 example.com\n0.0.0.0 bad.com",
			expectedValid:   []string{"example.com", "bad.com"},
			expectedInvalid: []string{"::1 localhost"},
		},
		{
			name:            "entries with comments",
			content:         "# This is a comment\n127.0.0.1 example.com\n! Another comment\n0.0.0.0 malicious.com\n// Yet another comment\n127.0.0.1 ads.com",
			expectedValid:   []string{"example.com", "malicious.com", "ads.com"},
			expectedInvalid: nil,
		},
		{
			name:            "entries with inline comments",
			content:         "127.0.0.1 example.com # inline comment\n0.0.0.0 malicious.com # another inline comment",
			expectedValid:   []string{"example.com", "malicious.com"},
			expectedInvalid: nil,
		},
		{
			name:            "entries with whitespace",
			content:         "  127.0.0.1   example.com  \n\t0.0.0.0\tmalicious.com\t\n  127.0.0.1  ads.com  ",
			expectedValid:   []string{"example.com", "malicious.com", "ads.com"},
			expectedInvalid: nil,
		},
		{
			name:          "invalid domain entries",
			content:       "127.0.0.1 invalid..domain\n0.0.0.0 .invalid.domain\n127.0.0.1 invalid-.domain",
			expectedValid: nil,
			expectedInvalid: []string{
				"127.0.0.1 invalid..domain",
				"0.0.0.0 .invalid.domain",
				"127.0.0.1 invalid-.domain",
			},
		},
		{
			name:            "mixed valid and invalid entries",
			content:         "127.0.0.1 example.com\n0.0.0.0 invalid..domain\n127.0.0.1 good.domain.com\n0.0.0.0 .bad.domain",
			expectedValid:   []string{"example.com", "good.domain.com"},
			expectedInvalid: []string{"0.0.0.0 invalid..domain", "0.0.0.0 .bad.domain"},
		},
		{
			name:            "empty lines and whitespace only",
			content:         "\n  \n\t\n   \n",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "lines with only IP address",
			content:         "127.0.0.1\n0.0.0.0",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "lines with extra fields",
			content:         "127.0.0.1 example.com extra field\n0.0.0.0 malicious.com another extra field",
			expectedValid:   []string{"example.com", "malicious.com"},
			expectedInvalid: nil,
		},
		{
			name:            "subdomain entries",
			content:         "127.0.0.1 sub.example.com\n0.0.0.0 deep.nested.sub.domain.com\n127.0.0.1 www.test.org",
			expectedValid:   []string{"sub.example.com", "deep.nested.sub.domain.com", "www.test.org"},
			expectedInvalid: nil,
		},
		{
			name:            "international domain names",
			content:         "127.0.0.1 example.co.uk\n0.0.0.0 test.org.au\n127.0.0.1 site.com.br",
			expectedValid:   []string{"example.co.uk", "test.org.au", "site.com.br"},
			expectedInvalid: nil,
		},
		{
			name:            "punycode domains",
			content:         "127.0.0.1 xn--example.com\n0.0.0.0 xn--test.org",
			expectedValid:   []string{"xn--example.com", "xn--test.org"},
			expectedInvalid: nil,
		},
		{
			name:            "localhost entries",
			content:         "127.0.0.1 localhost\n::1 localhost\n127.0.0.1 localhost.localdomain",
			expectedValid:   []string{"localhost.localdomain"},
			expectedInvalid: []string{"127.0.0.1 localhost", "::1 localhost"},
		},
		{
			name: "comprehensive hosts file format",
			content: `# Copyright (c) 1993-2009 Microsoft Corp.
#
# This is a sample HOSTS file used by Microsoft TCP/IP for Windows.
#
127.0.0.1       localhost
::1             localhost
127.0.0.1       example.com # blocked site
0.0.0.0         malicious.com
127.0.0.1       ads.tracker.com # ad tracker

# Some invalid entries
127.0.0.1       invalid..domain
0.0.0.0         .bad.start

# More valid entries
127.0.0.1       sub.domain.com
0.0.0.0         another.test.org`,
			expectedValid: []string{
				"example.com",
				"malicious.com",
				"ads.tracker.com",
				"sub.domain.com",
				"another.test.org",
			},
			expectedInvalid: []string{
				"127.0.0.1       localhost",
				"::1             localhost",
				"127.0.0.1       invalid..domain",
				"0.0.0.0         .bad.start",
			},
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

func TestHostnameProcessor_Integration(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewHostnameProcessor("hostname", "blocklist")

	content := `# This file contains the mappings of IP addresses to host names.
127.0.0.1       localhost
::1             localhost

# Block some ads and trackers
0.0.0.0         doubleclick.net
0.0.0.0         googleadservices.com

127.0.0.1       invalid..domain
0.0.0.0         .starting.with.dot
127.0.0.1       ending.with.dash-

127.0.0.1       sub.example.com
0.0.0.0         www.tracker.net`

	expectedValid := []string{
		"doubleclick.net",
		"googleadservices.com",
		"sub.example.com",
		"www.tracker.net",
	}

	expectedInvalid := []string{
		"127.0.0.1       localhost",
		"::1             localhost",
		"127.0.0.1       invalid..domain",
		"0.0.0.0         .starting.with.dot",
		"127.0.0.1       ending.with.dash-",
	}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid entries should match expected")
	assert.Equal(t, expectedInvalid, invalid, "Invalid entries should match expected")
}

func TestHostnameProcessor_EdgeCases(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewHostnameProcessor("hostname", "blocklist")

	tests := []struct {
		name            string
		content         string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "single line with no newline",
			content:         "127.0.0.1 example.com",
			expectedValid:   []string{"example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "tabs and mixed whitespace",
			content:         "127.0.0.1\t\texample.com\n0.0.0.0   \t  malicious.com",
			expectedValid:   []string{"example.com", "malicious.com"},
			expectedInvalid: nil,
		},
		{
			name:            "only comments",
			content:         "# Comment 1\n! Comment 2\n// Comment 3",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "malformed lines",
			content:         "notanip example.com\n127.0.0.1\njustadomain.com",
			expectedValid:   []string{"example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "very long domain",
			content:         "127.0.0.1 very.long.subdomain.name.that.might.be.valid.example.com",
			expectedValid:   []string{"very.long.subdomain.name.that.might.be.valid.example.com"},
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
