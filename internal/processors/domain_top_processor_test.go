package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewDomainTopProcessor(t *testing.T) {
	t.Parallel()

	sourceType := "domain_top"
	listType := "allowlist"

	processor := processors.NewDomainTopProcessor(sourceType, listType)

	assert.NotNil(t, processor, "Processor should not be nil")
	assert.Equal(t, sourceType, processor.GetSourceType(), "Source type should match")
	assert.Equal(t, listType, processor.GetListType(), "List type should match")
}

func TestDomainTopProcessor_Process(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainTopProcessor("domain_top", "allowlist")

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
			content:         "1, google.com",
			expectedValid:   []string{"google.com"},
			expectedInvalid: nil,
		},
		{
			name:            "multiple valid entries",
			content:         "1, google.com\n2, youtube.com\n3, facebook.com",
			expectedValid:   []string{"google.com", "youtube.com", "facebook.com"},
			expectedInvalid: nil,
		},
		{
			name:            "entries with subdomain",
			content:         "1, www.google.com\n2, mail.yahoo.com\n3, api.github.com",
			expectedValid:   []string{"www.google.com", "mail.yahoo.com", "api.github.com"},
			expectedInvalid: nil,
		},
		{
			name:            "mixed valid and invalid entries",
			content:         "1, google.com\ninvalid_line\n2, facebook.com\nbad format",
			expectedValid:   []string{"google.com", "facebook.com"},
			expectedInvalid: []string{"invalid_line", "bad format"},
		},
		{
			name:            "entries with different spacing",
			content:         "1,google.com\n2, youtube.com\n3,  facebook.com\n4,   twitter.com",
			expectedValid:   []string{"google.com", "youtube.com", "facebook.com", "twitter.com"},
			expectedInvalid: nil,
		},
		{
			name:          "invalid domain formats",
			content:       "1, invalid..domain\n2, .starts-with-dot\n3, ends-with-dot.\n4, -starts-with-dash",
			expectedValid: nil,
			expectedInvalid: []string{
				"1, invalid..domain",
				"2, .starts-with-dot",
				"3, ends-with-dot.",
				"4, -starts-with-dash",
			},
		},
		{
			name:            "comments and empty lines",
			content:         "# This is a comment\n1, google.com\n\n2, youtube.com\n! Another comment\n3, facebook.com",
			expectedValid:   []string{"google.com", "youtube.com", "facebook.com"},
			expectedInvalid: nil,
		},
		{
			name:            "large rank numbers",
			content:         "1000000, example.com\n999999, test.org\n1234567, sample.net",
			expectedValid:   []string{"example.com", "test.org", "sample.net"},
			expectedInvalid: nil,
		},
		{
			name:            "entries without rank",
			content:         "google.com\nyoutube.com\nfacebook.com",
			expectedValid:   nil,
			expectedInvalid: []string{"google.com", "youtube.com", "facebook.com"},
		},
		{
			name:            "entries with incorrect comma placement",
			content:         "1 google.com\n2,youtube,com\n3; facebook.com",
			expectedValid:   nil,
			expectedInvalid: []string{"1 google.com", "2,youtube,com", "3; facebook.com"},
		},
		{
			name:            "international domains",
			content:         "1, example.co.uk\n2, test.com.au\n3, sample.org.in\n4, site.edu.br",
			expectedValid:   []string{"example.co.uk", "test.com.au", "sample.org.in", "site.edu.br"},
			expectedInvalid: nil,
		},
		{
			name:            "domains with hyphens",
			content:         "1, my-domain.com\n2, test-site.org\n3, multi-word-domain.net",
			expectedValid:   []string{"my-domain.com", "test-site.org", "multi-word-domain.net"},
			expectedInvalid: nil,
		},
		{
			name:            "mixed case domains",
			content:         "1, Google.COM\n2, YouTUBE.com\n3, FaceBook.COM",
			expectedValid:   []string{"Google.COM", "YouTUBE.com", "FaceBook.COM"},
			expectedInvalid: nil,
		},
		{
			name:            "entries with trailing whitespace",
			content:         "1, google.com   \n2, youtube.com\t\n3, facebook.com \n",
			expectedValid:   []string{"google.com", "youtube.com", "facebook.com"},
			expectedInvalid: nil,
		},
		{
			name:            "only whitespace lines",
			content:         "   \n\t\n   \n",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "domains with numeric characters",
			content:         "1, 3m.com\n2, 7eleven.com\n3, 99designs.com",
			expectedValid:   []string{"3m.com", "7eleven.com", "99designs.com"},
			expectedInvalid: nil,
		},
		{
			name:            "entries with special characters in rank",
			content:         "1.0, google.com\n2.5, youtube.com\n3a, facebook.com",
			expectedValid:   nil,
			expectedInvalid: []string{"1.0, google.com", "2.5, youtube.com", "3a, facebook.com"},
		},
		{
			name:    "very long domain names",
			content: "1, very-long-domain-name-that-is-still-valid.com\n2, another-extremely-long-but-valid-domain-name.org",
			expectedValid: []string{
				"very-long-domain-name-that-is-still-valid.com",
				"another-extremely-long-but-valid-domain-name.org",
			},
			expectedInvalid: nil,
		},
		{
			name:            "domains with single character TLD",
			content:         "1, example.a\n2, test.x\n3, sample.z",
			expectedValid:   nil,
			expectedInvalid: []string{"1, example.a", "2, test.x", "3, sample.z"},
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

func TestDomainTopProcessor_Integration(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainTopProcessor("domain_top", "allowlist")

	content := `# Tranco Top Domains List
# Format: rank,domain
1, google.com
2, youtube.com
3, facebook.com
4, twitter.com
5, instagram.com
6, wikipedia.org
7, yahoo.com
8, whatsapp.com
invalid_entry_without_rank
10, invalid..domain
11, netflix.com
# End of list`

	expectedValid := []string{
		"google.com", "youtube.com", "facebook.com", "twitter.com", "instagram.com",
		"wikipedia.org", "yahoo.com", "whatsapp.com", "netflix.com",
	}
	expectedInvalid := []string{
		"invalid_entry_without_rank", "10, invalid..domain",
	}

	valid, invalid := processor.Process(logger, content)
	assert.Equal(t, expectedValid, valid, "Valid entries should match expected for realistic data")
	assert.Equal(t, expectedInvalid, invalid, "Invalid entries should match expected for realistic data")
}

func TestDomainTopProcessor_EdgeCases(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewDomainTopProcessor("domain_top", "allowlist")

	tests := []struct {
		name            string
		content         string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "only comments",
			content:         "# Comment 1\n! Comment 2\n// Comment 3",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "rank with leading zeros",
			content:         "001, google.com\n0002, youtube.com\n00003, facebook.com",
			expectedValid:   []string{"google.com", "youtube.com", "facebook.com"},
			expectedInvalid: nil,
		},
		{
			name:            "negative rank numbers",
			content:         "-1, google.com\n-2, youtube.com",
			expectedValid:   nil,
			expectedInvalid: []string{"-1, google.com", "-2, youtube.com"},
		},
		{
			name:            "rank with alphabetic characters",
			content:         "1a, google.com\nrank2, youtube.com\nthree, facebook.com",
			expectedValid:   nil,
			expectedInvalid: []string{"1a, google.com", "rank2, youtube.com", "three, facebook.com"},
		},
		{
			name:            "multiple commas",
			content:         "1,, google.com\n2, youtube, com\n3, face,book.com",
			expectedValid:   nil,
			expectedInvalid: []string{"1,, google.com", "2, youtube, com", "3, face,book.com"},
		},
		{
			name:            "domains starting with numbers",
			content:         "1, 123test.com\n2, 456-test.org\n3, 789.example.net",
			expectedValid:   []string{"123test.com", "456-test.org", "789.example.net"},
			expectedInvalid: nil,
		},
		{
			name:            "extremely large rank numbers",
			content:         "999999999999999999999, google.com\n1000000000000000000000, youtube.com",
			expectedValid:   []string{"google.com", "youtube.com"},
			expectedInvalid: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, invalid := processor.Process(logger, tt.content)
			assert.Equal(t, tt.expectedValid, valid, "Valid entries should match expected for edge case: %s", tt.name)
			assert.Equal(
				t,
				tt.expectedInvalid,
				invalid,
				"Invalid entries should match expected for edge case: %s",
				tt.name,
			)
		})
	}
}
