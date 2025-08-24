package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewIpv4FromDomainTopProcessor(t *testing.T) {
	sourceType := "ipv4_from_domain_top"
	listType := "allowlist"

	processor := processors.NewIpv4FromDomainProcessor(sourceType, listType)

	assert.NotNil(t, processor, "Processor should not be nil")
	assert.Equal(t, sourceType, processor.GetSourceType(), "Source type should match")
	assert.Equal(t, listType, processor.GetListType(), "List type should match")
}

func TestIpv4FromDomainProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewIpv4FromDomainProcessor("ipv4_from_domain", "allowlist")

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
			name:    "single valid entry",
			content: "example.com",
			expectedValid: []string{
				"23.192.228.80",
				"23.192.228.84",
				"23.215.0.136",
				"23.215.0.138",
				"23.220.75.232",
				"23.220.75.245",
			},
			expectedInvalid: nil,
		},
		{
			name:    "mixed valid and invalid entries",
			content: "example.com\ninvalid_line",
			expectedValid: []string{
				"23.192.228.80",
				"23.192.228.84",
				"23.215.0.136",
				"23.215.0.138",
				"23.220.75.232",
				"23.220.75.245",
			},
			expectedInvalid: []string{"invalid_line"},
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

func TestIpv4FromDomainProcessor_Integration(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewIpv4FromDomainProcessor("ipv4_from_domain", "allowlist")

	content := `# Domains List
# Format: domain
example.com
invalid_entry_without_rank
invalid..domain
# End of list`

	expectedValid := []string{
		"23.192.228.80",
		"23.192.228.84",
		"23.215.0.136",
		"23.215.0.138",
		"23.220.75.232",
		"23.220.75.245",
	}
	expectedInvalid := []string{
		"invalid..domain",
		"invalid_entry_without_rank",
	}

	valid, invalid := processor.Process(logger, content)
	assert.Equal(t, expectedValid, valid, "Valid entries should match expected for realistic data")
	assert.Equal(t, expectedInvalid, invalid, "Invalid entries should match expected for realistic data")
}

func TestIpv4FromDomainProcessor_Registration(t *testing.T) {
	_, existsBlock := processors.Processors.GetProcessor("ipv4_from_domain", "blocklist")
	_, existsAllow := processors.Processors.GetProcessor("ipv4_from_domain", "allowlist")

	assert.True(t, existsBlock, "ipv4_from_domain should be registered for blocklist")
	assert.True(t, existsAllow, "ipv4_from_domain should be registered for allowlist")
}
