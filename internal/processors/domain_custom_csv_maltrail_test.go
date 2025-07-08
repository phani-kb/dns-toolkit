package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewDomainCustomCsvMaltrailProcessor(t *testing.T) {
	t.Parallel()

	sourceType := "domain_custom_csv_maltrail"
	listType := "blocklist"

	processor := processors.NewDomainCustomCsvMaltrailProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestDomainCustomCsvMaltrailProcessor_Process(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvMaltrailProcessor("domain_custom_csv_maltrail", "blocklist")

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
			name:            "multiple valid domain entries",
			content:         "example.com,malware,bad site\nmalicious.org,phishing,phishing site\nads.tracker.com,tracker,ad tracker",
			expectedValid:   []string{"example.com", "malicious.org", "ads.tracker.com"},
			expectedInvalid: nil,
		},
		{
			name:            "domains with ports (invalid)",
			content:         "example.com:8080,malware,bad site\nmalicious.org:443,phishing,phishing site\ngood.domain.com,tracker,valid domain",
			expectedValid:   []string{"good.domain.com"},
			expectedInvalid: []string{"example.com:8080", "malicious.org:443"},
		},
		{
			name:            "invalid domain formats",
			content:         "invalid..domain,malware,bad domain\n.invalid.start,phishing,bad start\nvalid.domain.com,tracker,good domain",
			expectedValid:   []string{"valid.domain.com"},
			expectedInvalid: []string{"invalid..domain", ".invalid.start"},
		},
		{
			name:            "comments and file paths",
			content:         "# This is a comment\nexample.com,malware,description\n/path/to/file,file,skip this\nmalicious.org,phishing,phishing",
			expectedValid:   []string{"example.com", "malicious.org"},
			expectedInvalid: nil,
		},
		{
			name:            "real Maltrail sample format",
			content:         "priv.inteleksys.com,ek bottle (malicious),(static)\n4w6ylniamu6x7e3a.onion,ek bottle (malicious),(static)\nview.inteleksys.com,ek bottle (malicious),(static)\n150.158.31.113:8000,xiebroc2 (malicious),(static)\n150.158.31.113:8889,xiebroc2 (malicious),(static)",
			expectedValid:   []string{"priv.inteleksys.com", "4w6ylniamu6x7e3a.onion", "view.inteleksys.com"},
			expectedInvalid: []string{"150.158.31.113:8000", "150.158.31.113:8889"},
		},
		{
			name:            "mixed content with IP addresses",
			content:         "192.168.1.1,malware,ip address\nexample.com,phishing,valid domain\n10.0.0.1,tracker,another ip",
			expectedValid:   []string{"example.com"},
			expectedInvalid: []string{"192.168.1.1", "10.0.0.1"},
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

func TestDomainCustomCsvMaltrailProcessor_EdgeCases(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvMaltrailProcessor("domain_custom_csv_maltrail", "blocklist")

	tests := []struct {
		name            string
		content         string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "single line without newline",
			content:         "example.com,malware,description",
			expectedValid:   []string{"example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "only comments",
			content:         "# Comment 1\n! Comment 2\n// Comment 3",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "domains with underscores (invalid)",
			content:         "example-site.com,malware,hyphen domain\ntest_site.org,phishing,underscore domain\nsite123.net,tracker,numeric domain",
			expectedValid:   []string{"example-site.com", "site123.net"},
			expectedInvalid: []string{"test_site.org"},
		},
		{
			name:            "empty fields handling",
			content:         ",malware,empty domain\nexample.com,,valid domain\n,,,all empty\nmalicious.org,phishing,valid",
			expectedValid:   []string{"example.com", "malicious.org"},
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

func TestDomainCustomCsvMaltrailProcessor_Integration(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvMaltrailProcessor("domain_custom_csv_maltrail", "blocklist")

	// Real Maltrail sample with comments and mixed content
	content := `# Maltrail Domain Blacklist Export
priv.inteleksys.com,ek bottle (malicious),(static)
4w6ylniamu6x7e3a.onion,ek bottle (malicious),(static)
view.inteleksys.com,ek bottle (malicious),(static)
conforyou.ml,ek bottle (malicious),(static)
150.158.31.113:8000,xiebroc2 (malicious),(static)
150.158.31.113:8889,xiebroc2 (malicious),(static)
# Comment in the middle
invalid..domain.com,malware,(static)
.invalid.start.com,phishing,(static)
namelessserver.com,nameless c2 (malicious),(static)
192.168.1.100,malware,(static)`

	expectedValid := []string{
		"priv.inteleksys.com",
		"4w6ylniamu6x7e3a.onion",
		"view.inteleksys.com",
		"conforyou.ml",
		"namelessserver.com",
	}

	expectedInvalid := []string{
		"150.158.31.113:8000",
		"150.158.31.113:8889",
		"invalid..domain.com",
		".invalid.start.com",
		"192.168.1.100",
	}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid entries should match expected")
	assert.Equal(t, expectedInvalid, invalid, "Invalid entries should match expected")
}
