package main

import (
	"github.com/phani-kb/dns-toolkit/cmd"
)

// isHelpRequested checks if help is requested in the command line arguments
func isHelpRequested(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" || arg == "help" {
			return true
		}
	}
	return false
}

func main() {
	cmd.Execute()
}
