package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewDomainCustomCsvBlackbookProcessor(t *testing.T) {
	t.Parallel()

	sourceType := "domain_custom_csv_blackbook"
	listType := "blocklist"

	processor := processors.NewDomainCustomCsvBlackbookProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestDomainCustomCsvBlackbookProcessor_Process(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvBlackbookProcessor("domain_custom_csv_blackbook", "blocklist")

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
			name:    "blackbook format with spaces",
			content: "mock-win.av-example.com                                    #AVG         Activate Free Version\nmock-ping.av-example.com                                   #AVG         Install Freezes\nmock-static.av-example.net                                 #AVG         Website Login",
			expectedValid: []string{
				"mock-win.av-example.com",
				"mock-ping.av-example.com",
				"mock-static.av-example.net",
			},
			expectedInvalid: nil,
		},
		{
			name:            "with comments header",
			content:         "#####################################################\n# Antivirus\n#####################################################\nmock-av.example.com #AVG Activate Free Version\nmock-scan.example.com #Kaspersky Scans Webpages",
			expectedValid:   []string{"mock-av.example.com", "mock-scan.example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "invalid domains with ports",
			content:         "mock-update.example.com:8080 #McAfee Installer\nmock-valid.example.com #Norton Valid Domain\nmock-cloud.example.com:443 #Norton Check account",
			expectedValid:   []string{"mock-valid.example.com"},
			expectedInvalid: []string{"mock-update.example.com:8080", "mock-cloud.example.com:443"},
		},
		{
			name:            "invalid domain formats",
			content:         "invalid..domain #Software Bad domain\n.invalid.start #Software Bad start\nmock-valid.example.com #Software Good domain",
			expectedValid:   []string{"mock-valid.example.com"},
			expectedInvalid: []string{"invalid..domain", ".invalid.start"},
		},
		{
			name:            "mixed valid and invalid entries",
			content:         "mock-analytics.example.com #Norton Won't Login\n192.168.1.1 #Software IP Address\nmock-updates.example.org #Sophos Updates\ninvalid_domain #Software Invalid underscore",
			expectedValid:   []string{"mock-analytics.example.com", "mock-updates.example.org"},
			expectedInvalid: []string{"192.168.1.1", "invalid_domain"},
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

func TestDomainCustomCsvBlackbookProcessor_EdgeCases(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvBlackbookProcessor("domain_custom_csv_blackbook", "blocklist")

	tests := []struct {
		name            string
		content         string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "single line without newline",
			content:         "mock-single.example.com #Software Single entry",
			expectedValid:   []string{"mock-single.example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "only comments and headers",
			content:         "# Comment 1\n! Comment 2\n// Comment 3\n#####################################################",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "domains with special characters",
			content:         "mock-site.example.com #Software Hyphen domain\nmock_site.example.org #Software Underscore domain\nmock123.example.net #Software Numeric domain",
			expectedValid:   []string{"mock-site.example.com", "mock123.example.net"},
			expectedInvalid: []string{"mock_site.example.org"},
		},
		{
			name:    "vendor software blackbook format",
			content: "mock-dellupdater.example.com                                #Dell   Download \"Alienware Command Center\"\nmock-supportkb.example.com                                  #Dell   Support Page Images\nmock-www1.ins.example.com                                   #Dell   TechDirect - Parts Page\nmock-web-chat-e2ee.example.com                         #Facebook    Breaks 'To' in messages\nmock-mqtt-us-2.example.com                              #Roborock               Connect to Wifi / Check for firmware updates *Vacuum\nmock-media.dds.example.com                                #Lenovo   Media Creation Tool (Download)\nmock-tags.tiqcdn-example.com                             #Tracker - Breaks Download Link",
			expectedValid: []string{
				"mock-dellupdater.example.com",
				"mock-supportkb.example.com",
				"mock-www1.ins.example.com",
				"mock-web-chat-e2ee.example.com",
				"mock-mqtt-us-2.example.com",
				"mock-media.dds.example.com",
				"mock-tags.tiqcdn-example.com",
			},
			expectedInvalid: nil,
		},
		{
			name:            "international domains",
			content:         "mock-test.co.uk #Software UK domain\nmock-site.com.au #Software Australian domain\nmock-example.de #Software German domain",
			expectedValid:   []string{"mock-test.co.uk", "mock-site.com.au", "mock-example.de"},
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

func TestDomainCustomCsvBlackbookProcessor_AntivirusFormat(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvBlackbookProcessor("domain_custom_csv_blackbook", "blocklist")

	content := `#####################################################
# Antivirus
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

func TestDomainCustomCsvBlackbookProcessor_Integration(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvBlackbookProcessor("domain_custom_csv_blackbook", "blocklist")

	content := `#####################################################
# Software Blackbook Export
#####################################################
# Format: domain    #software    description
mock-telemetry.software-example.com #Software A Telemetry collection
mock-analytics.software-example.com #Software B Analytics tracking
# Invalid entries below
invalid..domain.com #Software C Invalid domain format
mock-ads.software-example.com:8080 #Software D Domain with port (invalid)
192.168.1.100 #Software E IP address (invalid)
.invalid.start.com #Software F Invalid start (invalid)

# Valid entries with special cases (blackbook format)
mock-hyphen-domain.software-example.com           #Software G    Domain with hyphens
mock123numeric.software-example.org               #Software H    Domain with numbers
mock_underscore.software-example.net              #Software I    Domain with underscore (invalid)

# Empty and whitespace handling
mock-whitespace.software-example.com   #Software K Domain with trailing space
   mock-leading.software-example.com #Software L Domain with leading space

# International domains (blackbook format)
mock-uk.software-example.co.uk #Software M UK domain
mock-de.software-example.de                       #Software N    German domain`

	expectedValid := []string{
		"mock-telemetry.software-example.com",
		"mock-analytics.software-example.com",
		"mock-hyphen-domain.software-example.com",
		"mock123numeric.software-example.org",
		"mock-whitespace.software-example.com",
		"mock-leading.software-example.com",
		"mock-uk.software-example.co.uk",
		"mock-de.software-example.de",
	}

	expectedInvalid := []string{
		"invalid..domain.com",
		"mock-ads.software-example.com:8080",
		"192.168.1.100",
		".invalid.start.com",
		"mock_underscore.software-example.net",
	}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid entries should match expected")
	assert.Equal(t, expectedInvalid, invalid, "Invalid entries should match expected")
}

func TestDomainCustomCsvBlackbookProcessor_RegistryIntegration(t *testing.T) {
	t.Parallel()

	processor, exists := processors.Processors.GetProcessor("domain_custom_csv_blackbook", "blocklist")
	assert.True(t, exists, "Processor should be registered")
	assert.NotNil(t, processor, "Processor should not be nil")

	blackbookProcessor, ok := processor.(*processors.DomainCustomCsvBlackbookProcessor)
	assert.True(t, ok, "Should be a DomainCustomCsvBlackbookProcessor")
	assert.Equal(t, "domain_custom_csv_blackbook", blackbookProcessor.GetSourceType())
	assert.Equal(t, "blocklist", blackbookProcessor.GetListType())
}
