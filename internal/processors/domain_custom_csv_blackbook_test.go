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
			name:            "basic CSV format with header",
			content:         "Domain,Malware,Date added,Source\nstatsrvv.com,keitaro,2024-09-01,ViriBack\n1312services.ru,1312,2024-09-01,ViriBack\nkaspersky-secure.ru,unam,2024-08-31,ViriBack",
			expectedValid:   []string{"statsrvv.com", "1312services.ru", "kaspersky-secure.ru"},
			expectedInvalid: nil,
		},
		{
			name:            "CSV with invalid domains",
			content:         "Domain,Malware,Date added,Source\ninvalid..domain.com,malware1,2024-09-01,Source1\nvalid-domain.com,malware2,2024-09-01,Source2\n192.168.1.1,malware3,2024-09-01,Source3",
			expectedValid:   []string{"valid-domain.com"},
			expectedInvalid: []string{"invalid..domain.com", "192.168.1.1"},
		},
		{
			name:            "CSV with empty lines and whitespace",
			content:         "Domain,Malware,Date added,Source\nwerdotx.shop,lokibot,2024-08-28,ViriBack\n\n   \nsodiumlaurethsulfatedesyroyer.com,lokibot,2024-08-18,ViriBack",
			expectedValid:   []string{"werdotx.shop", "sodiumlaurethsulfatedesyroyer.com"},
			expectedInvalid: nil,
		},
		{
			name:            "CSV with comments mixed in",
			content:         "Domain,Malware,Date added,Source\n# This is a comment\nsimple-updatereport.com,amadey,2024-08-13,ViriBack\n# Another comment\nruspyc.top,amadey,2024-08-11,ViriBack",
			expectedValid:   []string{"simple-updatereport.com", "ruspyc.top"},
			expectedInvalid: nil,
		},
		{
			name:            "CSV with domain containing port (invalid)",
			content:         "Domain,Malware,Date added,Source\nactualisation-service.com,amadey,2024-08-11,ViriBack\napi.garageserviceoperation.com:8080,amadey,2024-08-11,ViriBack\nsync.hiddenvnc.com,anonvnc,2024-08-11,ViriBack",
			expectedValid:   []string{"actualisation-service.com", "sync.hiddenvnc.com"},
			expectedInvalid: []string{"api.garageserviceoperation.com:8080"},
		},
		{
			name:    "CSV with various TLDs and subdomains",
			content: "Domain,Malware,Date added,Source\nsyn.hiddenvnc.com,anonvnc,2024-08-11,ViriBack\nsync.smartcloudflare.com,anonvnc,2024-08-11,ViriBack\nsync.smart-vnc.com,anonvnc,2024-08-11,ViriBack\ninvoice-traffic.com,anonvnc,2024-08-11,ViriBack",
			expectedValid: []string{
				"syn.hiddenvnc.com",
				"sync.smartcloudflare.com",
				"sync.smart-vnc.com",
				"invoice-traffic.com",
			},
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
			name:            "header only",
			content:         "Domain,Malware,Date added,Source",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "CSV with missing fields",
			content:         "Domain,Malware,Date added,Source\ntop.enkey.xyz\nsuburand.com,matanbuchus,2024-08-01,ViriBack",
			expectedValid:   []string{"top.enkey.xyz", "suburand.com"},
			expectedInvalid: nil,
		},
		{
			name:            "CSV with extra commas",
			content:         "Domain,Malware,Date added,Source\nrebistars.com,matanbuchus,2024-08-01,ViriBack,extra,field\nserak.top,lokibot,2024-07-31,ViriBack",
			expectedValid:   []string{"rebistars.com", "serak.top"},
			expectedInvalid: nil,
		},
		{
			name:            "domains with underscores (invalid)",
			content:         "Domain,Malware,Date added,Source\nvalid-domain.com,malware1,2024-09-01,Source1\ninvalid_domain.com,malware2,2024-09-01,Source2\nanother-valid.org,malware3,2024-09-01,Source3",
			expectedValid:   []string{"valid-domain.com", "another-valid.org"},
			expectedInvalid: []string{"invalid_domain.com"},
		},
		{
			name:            "domains starting with dots (invalid)",
			content:         "Domain,Malware,Date added,Source\n.invalid.start.com,malware1,2024-09-01,Source1\nvalid.domain.com,malware2,2024-09-01,Source2\n..double.dot.com,malware3,2024-09-01,Source3",
			expectedValid:   []string{"valid.domain.com"},
			expectedInvalid: []string{".invalid.start.com", "..double.dot.com"},
		},
		{
			name:            "CSV with quoted fields",
			content:         "Domain,Malware,Date added,Source\n\"quoted.domain.com\",\"malware with spaces\",\"2024-09-01\",\"Source Name\"\nnormal.domain.com,normalmalware,2024-09-01,NormalSource",
			expectedValid:   []string{"normal.domain.com"},
			expectedInvalid: []string{"\"quoted.domain.com\""},
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

func TestDomainCustomCsvBlackbookProcessor_ActualData(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvBlackbookProcessor("domain_custom_csv_blackbook", "blocklist")

	content := `Domain,Malware,Date added,Source
statsrvv.com,keitaro,2024-09-01,ViriBack
1312services.ru,1312,2024-09-01,ViriBack
kaspersky-secure.ru,unam,2024-08-31,ViriBack
werdotx.shop,lokibot,2024-08-28,ViriBack
sodiumlaurethsulfatedesyroyer.com,lokibot,2024-08-18,ViriBack
simple-updatereport.com,amadey,2024-08-13,ViriBack
ruspyc.top,amadey,2024-08-11,ViriBack
actualisation-service.com,amadey,2024-08-11,ViriBack
api.garageserviceoperation.com,amadey,2024-08-11,ViriBack
sync.hiddenvnc.com,anonvnc,2024-08-11,ViriBack
syn.hiddenvnc.com,anonvnc,2024-08-11,ViriBack
sync.smartcloudflare.com,anonvnc,2024-08-11,ViriBack
sync.smart-vnc.com,anonvnc,2024-08-11,ViriBack
invoice-traffic.com,anonvnc,2024-08-11,ViriBack
top.enkey.xyz,pony,2024-08-04,ViriBack
suburand.com,matanbuchus,2024-08-01,ViriBack
rebistars.com,matanbuchus,2024-08-01,ViriBack
serak.top,lokibot,2024-07-31,ViriBack`

	expectedValid := []string{
		"statsrvv.com",
		"1312services.ru",
		"kaspersky-secure.ru",
		"werdotx.shop",
		"sodiumlaurethsulfatedesyroyer.com",
		"simple-updatereport.com",
		"ruspyc.top",
		"actualisation-service.com",
		"api.garageserviceoperation.com",
		"sync.hiddenvnc.com",
		"syn.hiddenvnc.com",
		"sync.smartcloudflare.com",
		"sync.smart-vnc.com",
		"invoice-traffic.com",
		"top.enkey.xyz",
		"suburand.com",
		"rebistars.com",
		"serak.top",
	}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid domains should match expected from real-world data")
	assert.Empty(t, invalid, "No invalid entries expected for real-world CSV data")
}

func TestDomainCustomCsvBlackbookProcessor_MixedValidInvalid(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvBlackbookProcessor("domain_custom_csv_blackbook", "blocklist")

	content := `Domain,Malware,Date added,Source
# This is a comment line that should be skipped
valid-domain1.com,malware1,2024-09-01,Source1
invalid..domain.com,malware2,2024-09-01,Source2
192.168.1.100,malware3,2024-09-01,Source3
valid-domain2.org,malware4,2024-09-01,Source4
domain_with_underscore.net,malware5,2024-09-01,Source5
.starts.with.dot.com,malware6,2024-09-01,Source6
normal.domain.edu,malware7,2024-09-01,Source7
domain.with.port.com:8080,malware8,2024-09-01,Source8
another-valid.co.uk,malware9,2024-09-01,Source9

# Another comment
final.valid.domain.io,malware10,2024-09-01,Source10`

	expectedValid := []string{
		"valid-domain1.com",
		"valid-domain2.org",
		"normal.domain.edu",
		"another-valid.co.uk",
		"final.valid.domain.io",
	}

	expectedInvalid := []string{
		"invalid..domain.com",
		"192.168.1.100",
		"domain_with_underscore.net",
		".starts.with.dot.com",
		"domain.with.port.com:8080",
	}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid entries should match expected")
	assert.Equal(t, expectedInvalid, invalid, "Invalid entries should match expected")
}

func TestDomainCustomCsvBlackbookProcessor_Integration(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainCustomCsvBlackbookProcessor("domain_custom_csv_blackbook", "blocklist")

	content := `Domain,Malware,Date added,Source
# CSV format for malware domains
example1.com,trojan,2024-09-01,ThreatFeed
example2.org,ransomware,2024-08-30,SecurityVendor
# Invalid entries for testing
invalid..domain.net,virus,2024-08-29,Source3
192.168.1.1,botnet,2024-08-28,Source4
valid.domain.co.uk,malware,2024-08-27,Source5
.invalid.start.com,spyware,2024-08-26,Source6

# Whitespace and edge cases
  spaced.domain.com  ,adware,2024-08-25,Source7
domain_underscore.org,rootkit,2024-08-24,Source8
normal-hyphen.net,keylogger,2024-08-23,Source9

# International domains
german.example.de,trojan,2024-08-22,GermanSource
australian.example.com.au,virus,2024-08-21,AustralianSource`

	expectedValid := []string{
		"example1.com",
		"example2.org",
		"valid.domain.co.uk",
		"spaced.domain.com",
		"normal-hyphen.net",
		"german.example.de",
		"australian.example.com.au",
	}

	expectedInvalid := []string{
		"invalid..domain.net",
		"192.168.1.1",
		".invalid.start.com",
		"domain_underscore.org",
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
