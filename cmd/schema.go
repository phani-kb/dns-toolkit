package cmd

import (
	"context"
	"fmt"
	"os"
	"sort"

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
		embeddedChecksum := db.SchemaChecksum()

		fmt.Printf("Database path:     %s\n", dbPath)
		fmt.Printf("Embedded checksum: %s\n", embeddedChecksum)

		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			fmt.Printf("Stored checksum: <none>\n")
			fmt.Println("Schema: not initialized")
			os.Exit(0)
		}

		database, err := db.OpenInspect(dbPath)
		if err != nil {
			Logger.Errorf("Failed to open database: %v", err)
			os.Exit(1)
		}
		defer database.CloseLogError(Logger)
		storedChecksum := database.StoredChecksum(Logger)
		fmt.Printf("Stored checksum: %s\n", storedChecksum)

		switch storedChecksum {
		case "":
			fmt.Println("Schema: not initialized")
			os.Exit(0)
		case embeddedChecksum:
			fmt.Println("Schema: up to date")
		default:
			fmt.Println("Schema: out of date (will rebuild on next run)")
		}

		counts, err := database.TableRowCounts(Logger)
		if err != nil {
			Logger.Errorf("Failed to get table counts: %v", err)
			os.Exit(1)
		}

		if len(counts) > 0 {
			fmt.Println("\nTable row counts:")
			names := make([]string, 0, len(counts))
			for name := range counts {
				names = append(names, name)
			}
			sort.Strings(names)
			for _, name := range names {
				fmt.Printf("  %-35s %d\n", name, counts[name])
			}
		}
	},
}

var schemaHashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Print the SHA-256 checksum of the embedded schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(db.SchemaChecksum())
	},
}

var schemaRecreateCmd = &cobra.Command{
	Use:   "recreate",
	Short: "Force drop and recreate the database schema",
	Long: `Drops all tables and rebuilds the schema from schema.sql.
A full re-download and reprocess will be needed after this.`,
	Run: func(cmd *cobra.Command, args []string) {
		forceFlag, err := cmd.Flags().GetBool("force")
		if err != nil {
			Logger.Errorf("Failed to get force: %v", err)
		}
		forceEnv := os.Getenv("DNS_TOOLKIT_FORCE_RECREATE_DB") == constants.BooleanTrue

		if !forceFlag && !forceEnv {
			fmt.Println("Use --force to confirm schema recreation.")
			os.Exit(1)
		}

		ctx := context.Background()
		dbPath := getDBPath()
		database, err := db.Open(ctx, Logger, dbPath, true)
		if err != nil {
			Logger.Errorf("Failed to recreate schema: %v", err)
			os.Exit(1)
		}
		database.CloseLogError(Logger)

		fmt.Printf("Schema recreated at %s\n", dbPath)
		fmt.Printf("New Checksum: %s\n", db.SchemaChecksum())
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
