package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewIpv6HtaccessProcessor(t *testing.T) {
	sourceType := "ipv6_htaccess"
	listType := "blocklist"

	processor := processors.NewIpv6HtaccessProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestIpv6HtaccessProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewIpv6HtaccessProcessor("ipv6_htaccess", "blocklist")

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
			name: "htaccess content with mixed ipv4 and ipv6",
			content: `deny from 35.206.80.53
deny from 2001:9b1:28fa:cf00:cfbc:17d2:84bb:722e
deny from 2a01:111:2054:15f::aff:feaf:6f06
deny from 2402:1980:84cb:fb82::1`,
			expectedValid: []string{
				"2001:9b1:28fa:cf00:cfbc:17d2:84bb:722e",
			},
			expectedInvalid: []string{
				"deny from 35.206.80.53",
				"deny from 2a01:111:2054:15f::aff:feaf:6f06",
				"deny from 2402:1980:84cb:fb82::1",
			},
		},
		{
			name:            "content with full ipv6",
			content:         "deny from 2001:9b1:28fa:cf00:cfbc:17d2:84bb:722e",
			expectedValid:   []string{"2001:9b1:28fa:cf00:cfbc:17d2:84bb:722e"},
			expectedInvalid: nil,
		},
		{
			name:            "content with comments",
			content:         "# This is a comment\n# Another comment",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validEntries, invalidEntries := processor.Process(logger, tt.content)

			assert.Equal(t, tt.expectedValid, validEntries, "Valid entries should match")
			assert.Equal(t, tt.expectedInvalid, invalidEntries, "Invalid entries should match")
		})
	}
}

func TestIpv6HtaccessProcessor_ProcessorRegistration(t *testing.T) {
	processor, exists := processors.Processors.GetProcessor("ipv6_htaccess", "blocklist")
	assert.True(t, exists, "Processor should be registered")
	assert.NotNil(t, processor)

	ipv6HtaccessProcessor, ok := processor.(*processors.Ipv6HtaccessProcessor)
	assert.True(t, ok, "Processor should be of type *Ipv6HtaccessProcessor")
	assert.Equal(t, "ipv6_htaccess", ipv6HtaccessProcessor.GetSourceType())
	assert.Equal(t, "blocklist", ipv6HtaccessProcessor.GetListType())
}
