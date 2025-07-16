package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewDomainWithCommentSuffixProcessor(t *testing.T) {
	sourceType := "domain_with_comment_suffix"
	listType := "allowlist"

	processor := processors.NewDomainWithCommentSuffixProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestDomainWithCommentSuffixProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewDomainWithCommentSuffixProcessor("domain_with_comment_suffix", "allowlist")

	tests := []struct {
		name            string
		content         string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "with comments header",
			content:         "#####################################################\n# Antivirus\n#####################################################\nmock-av.example.com #AVG Activate Free Version\nmock-scan.example.com #Kaspersky Scans Webpages",
			expectedValid:   []string{"mock-av.example.com", "mock-scan.example.com"},
			expectedInvalid: nil,
		},
		{
			name:          "mixed valid and invalid entries",
			content:       "mock-analytics.example.com #Norton Won't Login\n192.168.1.1 #Software IP Address\nmock-updates.example.org #Sophos Updates\ninvalid_domain #Software Invalid underscore",
			expectedValid: []string{"mock-analytics.example.com", "mock-updates.example.org"},
			expectedInvalid: []string{
				"192.168.1.1 #Software IP Address",
				"invalid_domain #Software Invalid underscore",
			},
		},
		{
			name:            "empty lines and whitespace",
			content:         "mock-first.example.com #Software Description\n\n   \nmock-second.example.com #Software Another description\n   mock-third.example.com   #Software Trimmed",
			expectedValid:   []string{"mock-first.example.com", "mock-second.example.com", "mock-third.example.com"},
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

func TestDomainWithCommentSuffixProcessor_MixedFormat(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewDomainWithCommentSuffixProcessor("domain_with_comment_suffix", "allowlist")

	content := `#####################################################
# Mixed Format Example
#####################################################
mock-win.av-example.com                                    #AVG         Activate Free Version
mock-keepalive.security-example.com                        #McAfee      Installer (Internet Connection Test)
mock-provision.security-example.com                        #McAfee      Installer / Activation Key (Internet Connection Test)
mock-apps.security-example.com                             #McAfee      Installer (Errors out)
mock-analytics.security-example.com                        #Norton      Won't Login
mock-dellupdater.example.com                               #Dell        Download "Alienware Command Center"
mock-www1.ins.example.com                                  #Dell        TechDirect - Parts Page
mock-web-chat-e2ee.example.com                            #Facebook     Breaks 'To' in messages
mock-mqtt-us-2.example.com                                 #Roborock     Connect to Wifi / Check for firmware updates *Vacuum
mock-media.dds.example.com                                 #Lenovo      Media Creation Tool (Download)
mock-d17vo8z6jop21h.cloudfront-example.net                                 #Amazon - Music - Breaks it (App) - Audio
mock-ygf52iv00zfnxdocs4kfzcalhg.appsync-api.us-east-1.amazonaws-example.com #Amazon - Music - Breaks it (App) - Load Album Art
mock-tags.tiqcdn-example.com                               #Tracker     Breaks Download Link`

	expectedValid := []string{
		"mock-win.av-example.com",
		"mock-keepalive.security-example.com",
		"mock-provision.security-example.com",
		"mock-apps.security-example.com",
		"mock-analytics.security-example.com",
		"mock-dellupdater.example.com",
		"mock-www1.ins.example.com",
		"mock-web-chat-e2ee.example.com",
		"mock-mqtt-us-2.example.com",
		"mock-media.dds.example.com",
		"mock-d17vo8z6jop21h.cloudfront-example.net",
		"mock-ygf52iv00zfnxdocs4kfzcalhg.appsync-api.us-east-1.amazonaws-example.com",
		"mock-tags.tiqcdn-example.com",
	}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid antivirus domains should match expected")
	assert.Empty(t, invalid, "No invalid entries expected for valid antivirus format")
}

func TestDomainWithCommentSuffixProcessor_Integration(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewDomainWithCommentSuffixProcessor("domain_with_comment_suffix", "allowlist")

	content := `#####################################################
# Mixed Domain List Example
#####################################################
# Format: domain    #software    description
mock-telemetry.software-example.com #Software A Telemetry collection
mock-analytics.software-example.com #Software B Analytics tracking
# Invalid entries below
invalid..domain.com #Software C Invalid domain format
mock-ads.software-example.com:8080 #Software D Domain with port (invalid)
192.168.1.100 #Software E IP address (invalid)
.invalid.start.com #Software F Invalid start (invalid)

mock-hyphen-domain.software-example.com           #Software G    Domain with hyphens
mock123numeric.software-example.org               #Software H    Domain with numbers
mock_underscore.software-example.net              #Software I    Domain with underscore (invalid)

# Empty and whitespace handling
mock-whitespace.software-example.com   #Software K Domain with trailing space
   mock-leading.software-example.com #Software L Domain with leading space

mock-uk.software-example.co.uk #Software M UK domain
mock-de.software-example.de                       #Software N    German domain`

	expectedValid := []string{
		"mock-telemetry.software-example.com",
		"mock-analytics.software-example.com",
		"mock-hyphen-domain.software-example.com",
		"mock123numeric.software-example.org",
		"underscore.software-example.net", // Note: extracted from "mock_underscore..." (underscore stops the regex)
		"mock-whitespace.software-example.com",
		"mock-leading.software-example.com",
		"mock-uk.software-example.co.uk",
		"mock-de.software-example.de",
	}

	expectedInvalid := []string{
		"invalid..domain.com #Software C Invalid domain format",
		"mock-ads.software-example.com:8080 #Software D Domain with port (invalid)",
		"192.168.1.100 #Software E IP address (invalid)",
		".invalid.start.com #Software F Invalid start (invalid)",
	}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid entries should match expected")
	assert.Equal(t, expectedInvalid, invalid, "Invalid entries should match expected")
}

func TestDomainWithCommentSuffixProcessor_RegistryIntegration(t *testing.T) {
	processor, exists := processors.Processors.GetProcessor("domain_with_comment_suffix", "allowlist")
	assert.True(t, exists, "Processor should be registered")
	assert.NotNil(t, processor, "Processor should not be nil")

	domainWithSuffixProcessor, ok := processor.(*processors.DomainWithCommentSuffixProcessor)
	assert.True(t, ok, "Should be a DomainWithCommentSuffixProcessor")
	assert.Equal(t, "domain_with_comment_suffix", domainWithSuffixProcessor.GetSourceType())
	assert.Equal(t, "allowlist", domainWithSuffixProcessor.GetListType())
}
