package cmd

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractEntriesWithRegex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		content         string
		regexPattern    string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "extract domains",
			content:         "example.com\ntest.org\ninvalid-line\ngoogle.com\n",
			regexPattern:    `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`,
			expectedValid:   []string{"example.com", "test.org", "google.com"},
			expectedInvalid: []string{"invalid-line"},
		},
		{
			name:            "extract IPv4 addresses",
			content:         "192.168.1.1\n10.0.0.1\ninvalid-ip\n127.0.0.1\n",
			regexPattern:    `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`,
			expectedValid:   []string{"192.168.1.1", "10.0.0.1", "127.0.0.1"},
			expectedInvalid: []string{"invalid-ip"},
		},
		{
			name:            "empty content",
			content:         "",
			regexPattern:    `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`,
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "no matches",
			content:         "invalid-line1\ninvalid-line2\n",
			regexPattern:    `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`,
			expectedValid:   nil,
			expectedInvalid: []string{"invalid-line1", "invalid-line2"},
		},
		{
			name:            "all matches",
			content:         "example.com\ntest.org\n",
			regexPattern:    `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`,
			expectedValid:   []string{"example.com", "test.org"},
			expectedInvalid: nil,
		},
		{
			name:            "content with comments and empty lines",
			content:         "# Comment\nexample.com\n\ntest.org\n# Another comment\n",
			regexPattern:    `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`,
			expectedValid:   []string{"example.com", "test.org"}, // Comments and empty lines are filtered out
			expectedInvalid: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regex := regexp.MustCompile(tt.regexPattern)
			valid, invalid := extractEntriesWithRegex(tt.content, regex)

			if tt.expectedValid == nil {
				assert.Nil(t, valid, "Valid entries should be nil")
			} else {
				assert.ElementsMatch(t, tt.expectedValid, valid, "Valid entries should match")
			}

			if tt.expectedInvalid == nil {
				assert.Nil(t, invalid, "Invalid entries should be nil")
			} else {
				assert.ElementsMatch(t, tt.expectedInvalid, invalid, "Invalid entries should match")
			}
		})
	}
}
