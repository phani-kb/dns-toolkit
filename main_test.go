package main

import (
	"testing"
)

func TestIsHelpRequested(t *testing.T) {
	testCases := []struct {
		name        string
		args        []string
		expectsHelp bool
	}{
		{"help flag", []string{"dns-toolkit", "--help"}, true},
		{"short help flag", []string{"dns-toolkit", "-h"}, true},
		{"help command", []string{"dns-toolkit", "help"}, true},
		{"no help", []string{"dns-toolkit", "download"}, false},
		{"help with other args", []string{"dns-toolkit", "download", "--help"}, true},
		{"partial match should not trigger", []string{"dns-toolkit", "--helpful"}, false},
		{"empty args", []string{}, false},
		{"only program name", []string{"dns-toolkit"}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isHelpRequested(tc.args)
			if result != tc.expectsHelp {
				t.Errorf("Expected isHelpRequested(%v) = %v, got %v", tc.args, tc.expectsHelp, result)
			}
		})
	}
}
