package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewIpv4CsvHttpUrlFindProcessor(t *testing.T) {
	sourceType := "ipv4_csv_http_url_find"
	listType := "blocklist"

	processor := processors.NewIpv4CsvHttpUrlFindProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestIpv4CsvHttpUrlFindProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewIpv4CsvHttpUrlFindProcessor("ipv4_csv_http_url_find", "blocklist")

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
			name:            "CSV with domain in URL (should be invalid)",
			content:         "Keitaro,http://traffflo.pw/admin/,5.8.88.26,10-06-2019",
			expectedValid:   nil,
			expectedInvalid: []string{"Keitaro,http://traffflo.pw/admin/,5.8.88.26,10-06-2019"},
		},
		{
			name:            "CSV with IP in URL",
			content:         "Pony,http://185.79.156.18/t/d1/admin.php,185.79.156.18,11-06-2019",
			expectedValid:   []string{"185.79.156.18"},
			expectedInvalid: nil,
		},
		{
			name:            "mixed valid and invalid entries",
			content:         "ValidEntry,http://192.168.1.1/admin,192.168.1.1,10-06-2019\nInvalidEntry,no-url-here,not-an-ip,10-06-2019",
			expectedValid:   []string{"192.168.1.1"},
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

func TestIpv4CsvHttpUrlFindProcessor_Integration(t *testing.T) {
	logger := multilog.NewLogger()
	processor := processors.NewIpv4CsvHttpUrlFindProcessor("ipv4_csv_http_url_find", "blocklist")

	content := `Kpot,http://5.8.88.28/lBwKpCPQuLhfsuPU/login.php,5.8.88.28,10-06-2019
Pony,http://185.79.156.18/t/d1/admin.php,185.79.156.18,11-06-2019`

	expectedValid := []string{"5.8.88.28", "185.79.156.18"}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid entries should match expected")
	assert.Empty(t, invalid, "No invalid entries expected for this content")
}
