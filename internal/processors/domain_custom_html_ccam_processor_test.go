package processors

import (
	"testing"

	"github.com/phani-kb/multilog"
)

func TestDomainCustomHtmlCcamProcessor_Process(t *testing.T) {
	sampleContent := `<table><tr><td>Tayuya</td><td>13/06/2018 23:06:30</td><td>sisr.cacsite.com</td><td>fb541b4d571555c998341c0e856b8e5051f173a4</td></tr>
<tr><td>Tayuya</td><td>13/06/2018 23:03:31</td><td>growyourownteacher.co.uk</td><td>ef6ee6b8cb5c6e29cbd6888623f52c19aa13a4ff</td></tr>
<tr><td>Tayuya</td><td>13/06/2018 23:00:29</td><td>69.73.130.134</td><td>320e9a5d584b96d134c6ce7c541afb7794506a63</td></tr>
<tr><td>Tayuya</td><td>01/05/2018 20:34:52</td><td>185.99.133.162</td><td>44dfd1f712794fe4708e91d48af5ff78e728e221</td></tr>
<tr><td>Tayuya</td><td>01/05/2018 20:31:56</td><td>103.208.86.48</td><td>d058db07bcc20edae851d6aef8978bdb4d3df3fd</td></tr>
<tr><td>Tayuya</td><td>01/05/2018 20:28:47</td><td>aspmailcenter2.com</td><td>b168facd54cbe7b665677412f4780ebbb583bda2</td></tr>
<tr><td>Tayuya</td><td>01/05/2018 20:15:42</td><td>maxidoms.com</td><td>17b288e77bc73f4fe89c0e009384d1c0e794d1dc</td></tr>
<tr><td>Tayuya</td><td>01/05/2018 20:03:30</td><td>brausincsystem.pro</td><td>db58ba784b92dc64ff7e4a28078305de7117919d</td></tr></table>`

	logger := multilog.NewLogger()
	processor := NewDomainCustomHtmlCcamProcessor("domain_custom_html_ccam", "blocklist")
	valid, invalid := processor.Process(logger, sampleContent)

	expected := []string{
		"aspmailcenter2.com",
		"brausincsystem.pro",
		"growyourownteacher.co.uk",
		"maxidoms.com",
		"sisr.cacsite.com",
	}
	if len(valid) != len(expected) {
		t.Errorf("expected %d valid domains, got %d: %v", len(expected), len(valid), valid)
	}
	for i, domain := range expected {
		if i >= len(valid) || valid[i] != domain {
			t.Errorf("expected valid[%d]=%q, got %q", i, domain, valid[i])
		}
	}
	if len(invalid) == 0 {
		t.Error("expected some invalid entries, got none")
	}
}
