package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
)

func TestDomainUrlProcessor_Process(t *testing.T) {
	type testCase struct {
		name            string
		input           string
		expectedValid   []string
		expectedInvalid []string
	}

	tests := []testCase{
		{
			name:            "valid domain",
			input:           "example.com",
			expectedValid:   []string{"example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "valid domain with path",
			input:           "example.com/path",
			expectedValid:   []string{"example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "invalid line",
			input:           "not_a_domain",
			expectedValid:   nil,
			expectedInvalid: []string{"not_a_domain"},
		},
		{
			name:            "comment and blank",
			input:           "# comment\n\nexample.com",
			expectedValid:   []string{"example.com"},
			expectedInvalid: nil,
		},
		{
			name:            "mixed valid and invalid",
			input:           "example.com\ninvalid_domain\nanother.com/path",
			expectedValid:   []string{"example.com", "another.com"},
			expectedInvalid: []string{"invalid_domain"},
		},
		{
			name: "sample domain and ip urls",
			input: `etabline.xyz/asdfa/
voubucleonteri.xyz/adaafd/
goodlifestylenews.com/wp-admin/st.php
185.215.113.207/gb2pnjsjcs/login.php
favfa222v.xyz/fadfv/pasdfanel/affdmin.php
91.241.19.159/m7vvsw2dsQ/login.php
beasdgadi.ga/chffud/PvfffsaqDq929BSx_A_D_M1n_a.php`,
			expectedValid: []string{
				"etabline.xyz",
				"voubucleonteri.xyz",
				"goodlifestylenews.com",
				"favfa222v.xyz",
				"beasdgadi.ga",
			},
			expectedInvalid: []string{
				"185.215.113.207/gb2pnjsjcs/login.php",
				"91.241.19.159/m7vvsw2dsQ/login.php",
			},
		},
	}

	processor := processors.NewDomainUrlProcessor("domain_url", "block")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			valid, invalid := processor.Process(nil, tc.input)
			if len(valid) != len(tc.expectedValid) {
				t.Errorf("expected valid %v, got %v", tc.expectedValid, valid)
			}
			if len(invalid) != len(tc.expectedInvalid) {
				t.Errorf("expected invalid %v, got %v", tc.expectedInvalid, invalid)
			}
		})
	}
}
