package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	exactMatch         bool
	searchProcessed    bool
	searchConsolidated bool
	performDNSLookup   bool
	performCNAMELookup bool
	searchAguard       bool
)

var searchCmd = &cobra.Command{
	Use:   "search [domain or IP]",
	Short: "Search for a domain or IP in the processed files",
	Long:  `Search for a given domain among the valid processed domain files and report the sources in which it was found. Also looks for its IP address among the valid IPv4 processed files.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.ToLower(args[0])
		Logger.Infof("Searching for: %s", query)
	},
}

func init() {

	searchCmd.Flags().BoolVarP(&exactMatch, "exact", "e", false, "Perform exact match instead of substring match")
	searchCmd.Flags().BoolVarP(&searchProcessed, "processed", "p", true, "Search in processed files")
	searchCmd.Flags().BoolVarP(&searchConsolidated, "consolidated", "c", true, "Search in consolidated files")
	searchCmd.Flags().
		BoolVarP(&performDNSLookup, "dns", "d", false, "Perform DNS lookup for domain names to find associated IPs")
	searchCmd.Flags().
		BoolVarP(&performCNAMELookup, "cname", "n", true, "Perform CNAME lookup for domain names and search for CNAME records")
	searchCmd.Flags().
		BoolVarP(&searchAguard, "adguard", "g", false, "Search in AdGuard files")
}
