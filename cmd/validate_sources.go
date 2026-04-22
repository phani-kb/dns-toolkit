package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var validationPerformed bool

const validateSourcesCommandName = "validate-sources"

type schemaValidationConfig struct {
	DNSToolkit struct {
		SourceFiles []string `yaml:"source_files"`
	} `yaml:"dns_toolkit"`
}

func shouldPrintValidationProgress() bool {
	return len(os.Args) > 1 && os.Args[1] == validateSourcesCommandName
}

func printValidationProgress(format string, args ...any) {
	if shouldPrintValidationProgress() {
		fmt.Printf(format+"\n", args...)
	}
}

func resolvePathFromProjectRoot(projectRoot, path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(projectRoot, path)
}

func loadConfiguredSourceFiles(configPath string) ([]string, error) {
	projectRoot, err := utils.FindProjectRoot("")
	if err != nil {
		return nil, fmt.Errorf("finding project root: %w", err)
	}

	resolvedConfigPath := resolvePathFromProjectRoot(projectRoot, configPath)
	content, err := os.ReadFile(resolvedConfigPath)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var cfg schemaValidationConfig
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return nil, fmt.Errorf("decoding config file: %w", err)
	}

	return cfg.DNSToolkit.SourceFiles, nil
}

func validateSourcesSchema(configPath string) error {
	projectRoot, err := utils.FindProjectRoot("")
	if err != nil {
		return fmt.Errorf("finding project root: %w", err)
	}

	schemaPath := filepath.Join(projectRoot, constants.DataConfigDir, constants.SourcesSchemaFile)
	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(schemaPath)
	if err != nil {
		return fmt.Errorf("compiling sources schema: %w", err)
	}

	printValidationProgress("SCHEMA OK: %s", constants.SourcesSchemaFile)

	sourceFiles, err := loadConfiguredSourceFiles(configPath)
	if err != nil {
		return err
	}

	for _, sourceFile := range sourceFiles {
		resolvedSourceFile := resolvePathFromProjectRoot(projectRoot, sourceFile)

		if filepath.Base(resolvedSourceFile) == constants.SourcesSchemaFile {
			return fmt.Errorf("cannot validate schema file as input data: %s", sourceFile)
		}

		content, err := os.ReadFile(resolvedSourceFile)
		if err != nil {
			return fmt.Errorf("reading source file %s: %w", sourceFile, err)
		}

		if strings.TrimSpace(string(content)) == "" {
			Logger.Debugf("Skipping empty sources config file during schema validation: %s", sourceFile)
			printValidationProgress("SKIP (empty): %s", sourceFile)
			continue
		}

		var document any
		if err := json.Unmarshal(content, &document); err != nil {
			return fmt.Errorf("parsing source file %s: %w", sourceFile, err)
		}

		if err := schema.Validate(document); err != nil {
			return fmt.Errorf("schema validation failed for %s: %w", sourceFile, err)
		}

		printValidationProgress("PASS: %s", sourceFile)
	}

	return nil
}

func validateConfig(configPath string) error {
	if validationPerformed {
		slog.Debug("Skipping validation as it has already been performed")
		return nil
	}
	if configPath == "" {
		return fmt.Errorf("config validation error: %s", "config path is empty")
	}

	if err := validateSourcesSchema(configPath); err != nil {
		return fmt.Errorf("config validation error: %w", err)
	}

	appConfig, sourcesConfigs, err := config.LoadAppConfig(Logger, configPath)
	if err != nil {
		return fmt.Errorf("config validation error: %w", err)
	}

	AppConfig = &appConfig
	SourcesConfigs = sourcesConfigs
	validationPerformed = true

	return nil
}

var validateSourcesCmd = &cobra.Command{
	Use:   validateSourcesCommandName,
	Short: "Validate the sources configuration",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := GetConfigPath()
		if err != nil {
			slog.Error("Failed to get config path", "error", err)
			return
		}
		if err := validateConfig(configPath); err != nil {
			slog.Error("Validation failed", "error", err)
			return
		}
		Logger.Infof("Successfully loaded and validated configuration")
	},
}
