package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewIpv4CidrProcessor(t *testing.T) {
	sourceType := "ipv4_cidr_expand"
	listType := "blocklist"

	processor := processors.NewIpv4CidrProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestIpv4CidrProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewIpv4CidrProcessor("ipv4_cidr_expand", "blocklist")

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
			name:            "single valid CIDR /30",
			content:         "192.168.1.0/30",
			expectedValid:   []string{"192.168.1.0", "192.168.1.1", "192.168.1.2", "192.168.1.3"},
			expectedInvalid: nil,
		},
		{
			name:            "CIDR /31",
			content:         "10.0.0.0/31",
			expectedValid:   []string{"10.0.0.0", "10.0.0.1"},
			expectedInvalid: nil,
		},
		{
			name:            "single IP /32",
			content:         "172.16.0.1/32",
			expectedValid:   []string{"172.16.0.1"},
			expectedInvalid: nil,
		},
		{
			name:            "multiple CIDRs",
			content:         "192.168.1.0/30\n10.0.0.0/31",
			expectedValid:   []string{"192.168.1.0", "192.168.1.1", "192.168.1.2", "192.168.1.3", "10.0.0.0", "10.0.0.1"},
			expectedInvalid: nil,
		},
		{
			name:            "CIDR with comments",
			content:         "# Comment line\n192.168.1.0/30\n! Another comment",
			expectedValid:   []string{"192.168.1.0", "192.168.1.1", "192.168.1.2", "192.168.1.3"},
			expectedInvalid: nil,
		},
		{
			name:            "invalid CIDR format",
			content:         "192.168.1.0/33\n256.256.256.0/24",
			expectedValid:   nil,
			expectedInvalid: []string{"192.168.1.0/33", "256.256.256.0/24"},
		},
		{
			name:            "line without CIDR",
			content:         "not a cidr\n192.168.1.1\nrandom text",
			expectedValid:   nil,
			expectedInvalid: []string{"not a cidr", "192.168.1.1", "random text"},
		},
		{
			name:            "mixed valid and invalid",
			content:         "192.168.1.0/30\ninvalid line\n10.0.0.0/31",
			expectedValid:   []string{"192.168.1.0", "192.168.1.1", "192.168.1.2", "192.168.1.3", "10.0.0.0", "10.0.0.1"},
			expectedInvalid: []string{"invalid line"},
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

func TestIpv4CidrProcessor_Registration(t *testing.T) {
	processor, exists := processors.Processors.GetProcessor("ipv4_cidr_expand", "blocklist")
	assert.True(t, exists, "Processor should be registered")
	assert.NotNil(t, processor)

	ipv4CidrProcessor, ok := processor.(*processors.Ipv4CidrProcessor)
	assert.True(t, ok, "Processor should be of type *Ipv4CidrProcessor")
	assert.Equal(t, "ipv4_cidr_expand", ipv4CidrProcessor.GetSourceType())
	assert.Equal(t, "blocklist", ipv4CidrProcessor.GetListType())
}
