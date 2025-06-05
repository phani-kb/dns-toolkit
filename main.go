package main

import (
	"os"

	"github.com/phani-kb/dns-toolkit/cmd"
)

func main() {

	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" || arg == "help" {
			cmd.Execute()
			return
		}
	}
	cmd.Execute()
}
