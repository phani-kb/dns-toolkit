package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewIpv6FindProcessor(t *testing.T) {
	sourceType := "ipv6_find"
	listType := "blocklist"

	processor := processors.NewIpv6FindProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestIpv6FindProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewIpv6FindProcessor("ipv6_find", "blocklist")

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
			name:            "single valid ipv6",
			content:         "2001:470:1:c84::23",
			expectedValid:   []string{"2001:470:1:c84::23"},
			expectedInvalid: nil,
		},
		{
			name:            "ipv6 with additional data",
			content:         `2001:470:1:c84::23,"ftp,http,smtp,telnet"`,
			expectedValid:   []string{"2001:470:1:c84::23"},
			expectedInvalid: nil,
		},
		{
			name: "mixed content with ipv6 and ipv4",
			content: `2001:470:1:c84::23,"ftp,http,smtp,telnet"
61.6.225.90,smtp
2001:470:1:332::5,"ftp,http,smtp,telnet"`,
			expectedValid:   []string{"2001:470:1:c84::23", "2001:470:1:332::5"},
			expectedInvalid: []string{`61.6.225.90,smtp`},
		},
		{
			name:            "only ipv4 addresses",
			content:         "192.168.1.1",
			expectedValid:   nil,
			expectedInvalid: []string{"192.168.1.1"},
		},
		{
			name:            "content with comments",
			content:         "# This is a comment\n2001:470:1:c84::23",
			expectedValid:   []string{"2001:470:1:c84::23"},
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

func TestIpv6FindProcessor_ProcessorRegistration(t *testing.T) {
	processor, exists := processors.Processors.GetProcessor("ipv6_find", "blocklist")
	assert.True(t, exists, "Processor should be registered")
	assert.NotNil(t, processor)

	ipv6Processor, ok := processor.(*processors.Ipv6FindProcessor)
	assert.True(t, ok, "Processor should be of type *Ipv6FindProcessor")
	assert.Equal(t, "ipv6_find", ipv6Processor.GetSourceType())
	assert.Equal(t, "blocklist", ipv6Processor.GetListType())
}
