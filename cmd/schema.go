package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
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
		details, detailsErr := cmd.Flags().GetBool("details")
		if detailsErr != nil {
			Logger.Warnf("Failed to parse --details flag (defaulting to false): %v", detailsErr)
			details = false
		}

		dbPath := getDBPath()
		expectedChecksum := db.SchemaChecksum()

		fmt.Printf("Database path:     %s\n", dbPath)

		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			if details {
				fmt.Printf("Expected checksum: %s\n", expectedChecksum)
				fmt.Printf("Applied checksum:  <none>\n")
				fmt.Printf("Schema objects:    <none>\n")
			}
			fmt.Println("Schema: not initialized")
			os.Exit(0)
		}

		database, err := db.OpenInspect(dbPath)
		if err != nil {
			Logger.Errorf("Failed to open database: %v", err)
			os.Exit(1)
		}
		defer database.CloseLogError(Logger)
		appliedChecksum := database.StoredChecksum(Logger)

		if details {
			fmt.Printf("Expected checksum: %s\n", expectedChecksum)
			fmt.Printf("Applied checksum:  %s\n", appliedChecksum)
		}

		status := ""
		objectsMatch := false
		objectChecked := false

		switch appliedChecksum {
		case "":
			status = "not initialized"
			fmt.Printf("Schema: %s\n", status)
			os.Exit(0)
		case expectedChecksum:
			objectChecked = true
			liveObjectsChecksum, liveErr := database.LiveSchemaObjectsChecksum()
			if liveErr != nil {
				Logger.Errorf("Failed to compute live schema object checksum: %v", liveErr)
				os.Exit(1)
			}
			expectedObjectsChecksum := db.EmbeddedSchemaObjectsChecksum()
			objectsMatch = liveObjectsChecksum == expectedObjectsChecksum

			if objectsMatch {
				status = "up to date"
			} else {
				status = "drift detected (table/index objects differ from embedded schema)"
			}
			fmt.Printf("Schema: %s\n", status)
		default:
			status = "out of date (will rebuild on next run)"
			fmt.Printf("Schema: %s\n", status)
		}

		if details {
			switch {
			case !objectChecked:
				fmt.Println("Schema objects: not checked (schema checksum already out of date)")
			case objectsMatch:
				fmt.Println("Schema objects: OK")
			default:
				fmt.Println("Schema objects: DRIFT")
			}
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
	testMode := isTestMode()

	if AppConfig != nil && AppConfig.DNSToolkit.Database.Path != "" {
		dbPath := AppConfig.DNSToolkit.Database.Path
		if testMode {
			dbPath = remapDataDirForTests(dbPath)
		}
		if !filepath.IsAbs(dbPath) {
			if projectRoot := resolveProjectRoot(); projectRoot != "" {
				dbPath = filepath.Join(projectRoot, dbPath)
			}
		}
		return dbPath
	}

	if !testMode {
		return constants.DefaultDBPath
	}

	dbPath := constants.DefaultTestDBPath
	if !filepath.IsAbs(dbPath) {
		if projectRoot := resolveProjectRoot(); projectRoot != "" {
			dbPath = filepath.Join(projectRoot, dbPath)
		}
	}

	return dbPath
}

func init() {
	schemaStatusCmd.Flags().Bool("details", false, "Show checksum details")

	schemaRecreateCmd.Flags().Bool("force", false,
		"Confirm schema recreation (env: DNS_TOOLKIT_FORCE_RECREATE_DB)")

	schemaCmd.AddCommand(schemaStatusCmd)
	schemaCmd.AddCommand(schemaHashCmd)
	schemaCmd.AddCommand(schemaRecreateCmd)
}
