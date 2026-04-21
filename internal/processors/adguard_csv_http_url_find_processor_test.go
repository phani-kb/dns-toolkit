package processors

import (
	"testing"

	"github.com/phani-kb/multilog"
)

func TestAdguardCsvHttpUrlFindProcessor_Process(t *testing.T) {
	logger := multilog.NewLogger()
	processor := NewAdguardCsvHttpUrlFindProcessor("adguard_csv_http_url_find", "blocklist")

	// minimal, synthetic test CSV content (no real URLs)
	content := `A,http://host1.test/path,192.0.2.10,2025-10-05
B,https://host2.test/other,192.0.2.11,2025-10-05
C,http://192.0.2.1:8080/some,192.0.2.1,2025-10-05
`

	want := []string{
		"||host1.test/path$",
		"||host2.test/other$",
		"||192.0.2.1:8080/some$",
	}

	got, gotInv := processor.Process(logger, content)

	if len(gotInv) != 0 {
		t.Fatalf("expected no invalid lines, got: %v", gotInv)
	}

	if !EqualUnorderedAdg(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
