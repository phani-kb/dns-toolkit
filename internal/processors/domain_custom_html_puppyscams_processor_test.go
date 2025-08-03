package processors

import (
	"testing"

	"github.com/phani-kb/multilog"
)

func TestDomainCustomHtmlPuppyScamsProcessor_Process(t *testing.T) {
	sampleContent := `<a>example.com</a><a>fake-pets.net</a><a>invalid text</a>`

	logger := multilog.NewLogger()
	processor := NewDomainCustomHtmlPuppyScamsProcessor("domain_custom_html_puppyscams", "blocklist")
	valid, invalid := processor.Process(logger, sampleContent)

	expectedValid := []string{"example.com", "fake-pets.net"}
	if len(valid) != len(expectedValid) {
		t.Errorf("expected %d valid domains, got %d: %v", len(expectedValid), len(valid), valid)
	}

	for i, domain := range expectedValid {
		if i >= len(valid) || valid[i] != domain {
			t.Errorf("expected valid[%d]=%q, got %q", i, domain, valid[i])
		}
	}

	if len(invalid) != 1 || invalid[0] != "invalid text" {
		t.Errorf("expected 1 invalid entry 'invalid text', got %d: %v", len(invalid), invalid)
	}
}

func TestDomainCustomHtmlPuppyScamsProcessor_ProcessEmpty(t *testing.T) {
	logger := multilog.NewLogger()
	processor := NewDomainCustomHtmlPuppyScamsProcessor("domain_custom_html_puppyscams", "blocklist")
	valid, invalid := processor.Process(logger, "")

	if len(valid) != 0 || len(invalid) != 0 {
		t.Errorf("expected empty results for empty content, got valid: %v, invalid: %v", valid, invalid)
	}
}
