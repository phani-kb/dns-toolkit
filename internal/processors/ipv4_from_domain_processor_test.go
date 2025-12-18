package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIpv4FromDomainTopProcessor(t *testing.T) {
	sourceType := "ipv4_from_domain_top"
	listType := "allowlist"

	processor := processors.NewIpv4FromDomainProcessor(sourceType, listType)

	assert.NotNil(t, processor, "Processor should not be nil")
	assert.Equal(t, sourceType, processor.GetSourceType(), "Source type should match")
	assert.Equal(t, listType, processor.GetListType(), "List type should match")
}

func TestIpv4FromDomainProcessorProcess(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewIpv4FromDomainProcessor("ipv4_from_domain", "allowlist")

	tests := []struct {
		name               string
		content            string
		expectedValidLen   int
		expectedInvalidLen int
		invalidEntries     []string
	}{
		{
			name:               "empty content",
			content:            "",
			expectedValidLen:   0,
			expectedInvalidLen: 0,
			invalidEntries:     nil,
		},
		{
			name:               "single valid entry",
			content:            "example.com",
			expectedValidLen:   1,
			expectedInvalidLen: 0,
			invalidEntries:     nil,
		},
		{
			name:               "mixed valid and invalid entries",
			content:            "example.com\ninvalid_line",
			expectedValidLen:   1,
			expectedInvalidLen: 1,
			invalidEntries:     []string{"invalid_line"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, invalid := processor.Process(logger, tt.content)
			assert.GreaterOrEqual(t, len(valid), tt.expectedValidLen, "Valid entries count should be at least expected")
			assert.Equal(t, tt.expectedInvalidLen, len(invalid), "Invalid entries count should match expected")

			for _, ip := range valid {
				assert.True(t, utils.IsIPv4(ip), "Valid entry should be a valid IPv4 address: %s", ip)
			}

			if tt.invalidEntries != nil {
				assert.Equal(t, tt.invalidEntries, invalid, "Invalid entries should match expected")
			}
		})
	}
}

func TestIpv4FromDomainProcessorIntegration(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewIpv4FromDomainProcessor("ipv4_from_domain", "allowlist")

	content := `# Domains List
# Format: domain
example.com
invalid_entry_without_rank
invalid..domain
# End of list`

	valid, invalid := processor.Process(logger, content)

	require.GreaterOrEqual(t, len(valid), 1, "Should resolve at least one IP from example.com")

	for _, ip := range valid {
		assert.True(t, utils.IsIPv4(ip), "Valid entry should be a valid IPv4 address: %s", ip)
	}

	expectedInvalidCount := 2
	assert.Equal(t, expectedInvalidCount, len(invalid), "Should have %d invalid entries", expectedInvalidCount)
	assert.Contains(t, invalid, "invalid..domain", "Invalid domain format should be in invalid entries")
}

func TestIpv4FromDomainProcessorRegistration(t *testing.T) {
	_, existsBlock := processors.Processors.GetProcessor("ipv4_from_domain", "blocklist")
	_, existsAllow := processors.Processors.GetProcessor("ipv4_from_domain", "allowlist")

	assert.True(t, existsBlock, "ipv4_from_domain should be registered for blocklist")
	assert.True(t, existsAllow, "ipv4_from_domain should be registered for allowlist")
}
