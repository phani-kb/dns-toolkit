package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewAdGuardBlocklistProcessor(t *testing.T) {
	t.Parallel()

	sourceType := "adguard"
	listType := "blocklist"

	processor := processors.NewAdGuardBlocklistProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestAdGuardBlocklistProcessor_Process(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewAdGuardBlocklistProcessor("adguard", "blocklist")

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
			name:            "single valid entry",
			content:         "example.com",
			expectedValid:   []string{"example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "single exception entry",
			content:         "@@example.com",
			expectedValid:   nil,
			expectedInvalid: []string{"@@example.com"},
		},
		{
			name:            "mixed entries",
			content:         "example.com\n@@exception.com\nbad-site.com",
			expectedValid:   []string{"example.com", "bad-site.com"},
			expectedInvalid: []string{"@@exception.com"},
		},
		{
			name:            "entries with comments",
			content:         "# This is a comment\nexample.com\n! Another comment\n@@exception.com\n// Yet another comment\nbad-site.com",
			expectedValid:   []string{"example.com", "bad-site.com"},
			expectedInvalid: []string{"@@exception.com"},
		},
		{
			name:            "entries with whitespace",
			content:         "  example.com  \n\t@@exception.com\t\n  bad-site.com  ",
			expectedValid:   []string{"example.com", "bad-site.com"},
			expectedInvalid: []string{"@@exception.com"},
		},
		{
			name:            "adguard block format",
			content:         "||example.com^\n@@||exception.com^\n||bad-site.com^",
			expectedValid:   []string{"||example.com^", "||bad-site.com^"},
			expectedInvalid: []string{"@@||exception.com^"},
		},
		{
			name:            "empty lines and whitespace only",
			content:         "\n  \n\t\n   \n",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "complex adguard rules",
			content:         "||ads.example.com^\n@@||whitelist.example.com^\n! Comment\n# Another comment\n||tracking.example.com^$third-party\n@@||allowed.example.com^$important",
			expectedValid:   []string{"||ads.example.com^", "||tracking.example.com^$third-party"},
			expectedInvalid: []string{"@@||whitelist.example.com^", "@@||allowed.example.com^$important"},
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

func TestNewAdGuardAllowlistProcessor(t *testing.T) {
	t.Parallel()

	sourceType := "adguard"
	listType := "allowlist"

	processor := processors.NewAdGuardAllowlistProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestAdGuardAllowlistProcessor_Process(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewAdGuardAllowlistProcessor("adguard", "allowlist")

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
			name:            "single exception entry",
			content:         "@@example.com",
			expectedValid:   []string{"@@example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "single non-exception entry",
			content:         "example.com",
			expectedValid:   nil,
			expectedInvalid: []string{"example.com"},
		},
		{
			name:            "mixed entries",
			content:         "example.com\n@@exception.com\nbad-site.com",
			expectedValid:   []string{"@@exception.com"},
			expectedInvalid: []string{"example.com", "bad-site.com"},
		},
		{
			name:            "entries with comments",
			content:         "# This is a comment\nexample.com\n! Another comment\n@@exception.com\n// Yet another comment\nbad-site.com",
			expectedValid:   []string{"@@exception.com"},
			expectedInvalid: []string{"example.com", "bad-site.com"},
		},
		{
			name:            "entries with whitespace",
			content:         "  example.com  \n\t@@exception.com\t\n  bad-site.com  ",
			expectedValid:   []string{"@@exception.com"},
			expectedInvalid: []string{"example.com", "bad-site.com"},
		},
		{
			name:            "adguard allowlist format",
			content:         "||example.com^\n@@||exception.com^\n||bad-site.com^",
			expectedValid:   []string{"@@||exception.com^"},
			expectedInvalid: []string{"||example.com^", "||bad-site.com^"},
		},
		{
			name:            "empty lines and whitespace only",
			content:         "\n  \n\t\n   \n",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "complex adguard allowlist rules",
			content:         "||ads.example.com^\n@@||whitelist.example.com^\n! Comment\n# Another comment\n||tracking.example.com^$third-party\n@@||allowed.example.com^$important",
			expectedValid:   []string{"@@||whitelist.example.com^", "@@||allowed.example.com^$important"},
			expectedInvalid: []string{"||ads.example.com^", "||tracking.example.com^$third-party"},
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

func TestAdGuardProcessors_Integration(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	blocklistProcessor := processors.NewAdGuardBlocklistProcessor("adguard", "blocklist")
	allowlistProcessor := processors.NewAdGuardAllowlistProcessor("adguard", "allowlist")

	content := `# AdGuard Home blocklist
||ads.example.com^
@@||whitelist.example.com^
||tracking.site.com^$third-party
@@||allowed.site.com^$important
! This is a comment
example.domain.com
@@exception.domain.com`

	blocklistValid, blocklistInvalid := blocklistProcessor.Process(logger, content)
	allowlistValid, allowlistInvalid := allowlistProcessor.Process(logger, content)

	expectedBlocklistValid := []string{
		"||ads.example.com^",
		"||tracking.site.com^$third-party",
		"example.domain.com",
	}
	expectedBlocklistInvalid := []string{
		"@@||whitelist.example.com^",
		"@@||allowed.site.com^$important",
		"@@exception.domain.com",
	}

	expectedAllowlistValid := []string{
		"@@||whitelist.example.com^",
		"@@||allowed.site.com^$important",
		"@@exception.domain.com",
	}
	expectedAllowlistInvalid := []string{
		"||ads.example.com^",
		"||tracking.site.com^$third-party",
		"example.domain.com",
	}

	assert.Equal(t, expectedBlocklistValid, blocklistValid, "Blocklist valid entries should match")
	assert.Equal(t, expectedBlocklistInvalid, blocklistInvalid, "Blocklist invalid entries should match")
	assert.Equal(t, expectedAllowlistValid, allowlistValid, "Allowlist valid entries should match")
	assert.Equal(t, expectedAllowlistInvalid, allowlistInvalid, "Allowlist invalid entries should match")
}
