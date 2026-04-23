package cmd

import (
	"fmt"
	"os"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Manage the database schema",
	Long:  "Inspect, verify, or force-recreate the SQLite database schema.",
}

var schemaStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show schema status and table row counts",
	Run: func(cmd *cobra.Command, args []string) {
		dbPath := getDBPath()
		database, err := db.Open(dbPath)
		if err != nil {
			Logger.Errorf("Failed to open database: %v", err)
			os.Exit(1)
		}
		defer database.CloseLogError(Logger)
	},
}

var schemaHashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Print the SHA-256 checksum of the embedded schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO")
	},
}

var schemaRecreateCmd = &cobra.Command{
	Use:   "recreate",
	Short: "Force drop and recreate the database schema",
	Long: `Drops all tables and rebuilds the schema from schema.sql.
A full re-download and reprocess will be needed after this.`,
	Run: func(cmd *cobra.Command, args []string) {
		forceFlag, _ := cmd.Flags().GetBool("force")
		forceEnv := os.Getenv("DNS_TOOLKIT_FORCE_RECREATE_DB") == "true" ||
			os.Getenv("DNS_TOOLKIT_FORCE_RECREATE_DB") == "1"

		if !forceFlag && !forceEnv {
			fmt.Println("Use --force to confirm schema recreation.")
			os.Exit(1)
		}
	},
}

func getDBPath() string {
	if AppConfig != nil && AppConfig.DNSToolkit.Database.Path != "" {
		return AppConfig.DNSToolkit.Database.Path
	}
	return constants.DefaultDBPath
}

func init() {
	schemaRecreateCmd.Flags().Bool("force", false,
		"Confirm schema recreation (env: DNS_TOOLKIT_FORCE_RECREATE_DB)")

	schemaCmd.AddCommand(schemaStatusCmd)
	schemaCmd.AddCommand(schemaHashCmd)
	schemaCmd.AddCommand(schemaRecreateCmd)
}
