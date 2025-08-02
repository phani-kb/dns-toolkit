package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewDomainCsvHttpUrlFindProcessor(t *testing.T) {
	sourceType := "domain_csv_http_url_find"
	listType := "blocklist"

	processor := processors.NewDomainCsvHttpUrlFindProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestDomainCsvHttpUrlFindProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewDomainCsvHttpUrlFindProcessor("domain_csv_http_url_find", "blocklist")

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
			name:            "CSV with domain in URL",
			content:         "Keitaro,http://traffflo.pw/admin/,5.8.88.26,10-06-2019",
			expectedValid:   []string{"traffflo.pw"},
			expectedInvalid: nil,
		},
		{
			name:            "CSV with IP in URL (should not extract domain)",
			content:         "Pony,http://185.79.156.18/t/d1/admin.php,185.79.156.18,11-06-2019",
			expectedValid:   nil,
			expectedInvalid: []string{"Pony,http://185.79.156.18/t/d1/admin.php,185.79.156.18,11-06-2019"},
		},
		{
			name:            "mixed valid and invalid entries",
			content:         "ValidEntry,http://malicious.com/admin,192.168.1.1,10-06-2019\nInvalidEntry,no-url-here,not-an-ip,10-06-2019",
			expectedValid:   []string{"malicious.com"},
			expectedInvalid: []string{"InvalidEntry,no-url-here,not-an-ip,10-06-2019"},
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

func TestDomainCsvHttpUrlFindProcessor_Integration(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewDomainCsvHttpUrlFindProcessor("domain_csv_http_url_find", "blocklist")

	content := `Keitaro,http://traffflo.pw/admin/,5.8.88.26,10-06-2019
Kpot,http://benten02.futbol/QU6M6L2o04P9gIbD/login.php,8.209.74.36,11-06-2019`

	expectedValid := []string{"traffflo.pw", "benten02.futbol"}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid entries should match expected")
	assert.Empty(t, invalid, "No invalid entries expected for this content")
}

func TestDomainCsvHttpUrlFindProcessor_EdgeCases(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewDomainCsvHttpUrlFindProcessor("domain_csv_http_url_find", "blocklist")

	tests := []struct {
		name            string
		content         string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "URL with trailing dot",
			content:         "Malware,http://evil-site.com./admin,1.1.1.1,10-06-2019",
			expectedValid:   []string{"evil-site.com"},
			expectedInvalid: nil,
		},
		{
			name:            "mixed IP and domain URLs",
			content:         "IPUrl,http://192.168.1.1/admin,1.1.1.1,10-06-2019\nDomainUrl,http://evil.com/panel,2.2.2.2,11-06-2019",
			expectedValid:   []string{"evil.com"},
			expectedInvalid: []string{"IPUrl,http://192.168.1.1/admin,1.1.1.1,10-06-2019"},
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
