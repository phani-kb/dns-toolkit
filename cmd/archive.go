package cmd

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive DNS toolkit data",
	Long: `Archive DNS toolkit data from various folders specified in the config.
Creates a compressed archive (tgz) containing all relevant data files and
generates a summary of the archived contents.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := u.EnsureDirectoryExists(Logger, constants.ArchiveDir); err != nil {
			Logger.Errorf("Failed to create archive directory: %v", err)
			os.Exit(1)
		}

		runArchive(Logger)
	},
}

func runArchive(logger *multilog.Logger) {
	logger.Infof("Starting archive process")

	// Create a timestamp for archive naming
	timestamp := time.Now().Format(constants.TimestampFormat)
	archiveFilename := fmt.Sprintf("dns_toolkit_archive_%s.tgz", timestamp)
	archivePath := filepath.Join(constants.ArchiveDir, archiveFilename)

	// Ensure the archive directory exists
	if err := os.MkdirAll(constants.ArchiveDir, 0755); err != nil {
		logger.Errorf("Failed to create archive directory: %v", err)
		os.Exit(1)
	}

	archiveSummary := &common.ArchiveSummary{
		Folders:      []common.ArchiveFolder{},
		SummaryFiles: []common.ArchiveSummaryFile{},
	}

	archiveFile, err := os.Create(archivePath)
	if err != nil {
		logger.Errorf("Failed to create archive file: %v", err)
		os.Exit(1)
	}
	defer func() {
		if err := archiveFile.Close(); err != nil {
			logger.Warnf("Failed to close archive file: %v", err)
		}
	}()

	gzipWriter := gzip.NewWriter(archiveFile)
	defer func() {
		if err := gzipWriter.Close(); err != nil {
			logger.Warnf("Failed to close gzip writer: %v", err)
		}
	}()

	tarWriter := tar.NewWriter(gzipWriter)
	defer func() {
		if err := tarWriter.Close(); err != nil {
			logger.Warnf("Failed to close tar writer: %v", err)
		}
	}()

	summaryDir := AppConfig.DNSToolkit.Folders.Summary
	processSummaryFiles(logger, summaryDir, archiveSummary, tarWriter)

	foldersToArchive := u.GetFoldersToArchive(logger, constants.Folders)

	for folderPath, folderType := range foldersToArchive {
		logger.Infof("Processing folder: %s", folderPath)
		archiveFolder := common.ArchiveFolder{
			Name:      folderType + "(" + filepath.Base(folderPath) + ")",
			Files:     []common.ArchiveFile{},
			Timestamp: timestamp,
		}
		err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if path == folderPath {
				return nil
			}

			parent := filepath.Dir(path)

			// If the parent is not the folder we're directly processing, skip it (it's in a subdirectory)
			if parent != folderPath {
				if d.IsDir() {
					// Skip this directory and all its children
					return fs.SkipDir
				}
				return nil
			}

			// Only process files, not directories
			if !d.IsDir() {
				fileInfo, err := os.Stat(path)
				if err != nil {
					logger.Warnf("Failed to get file info for %s: %v", path, err)
					return nil
				}

				checksum := u.CalculateChecksum(
					logger,
					path,
					AppConfig.DNSToolkit.FilesChecksum.Algorithm,
				)

				archiveFile := common.ArchiveFile{
					Name:      filepath.Base(path),
					Filepath:  path,
					Checksum:  checksum,
					Size:      fileInfo.Size(),
					Timestamp: timestamp,
				}

				archiveFolder.Files = append(archiveFolder.Files, archiveFile)

				targetPath := filepath.Join(filepath.Base(folderPath), filepath.Base(path))
				if constants.OutputDir == folderPath {
					targetPath = filepath.Base(path)
				}

				err = addFileToTar(logger, tarWriter, path, targetPath, fileInfo)
				if err != nil {
					logger.Warnf("Failed to add file %s to archive: %v", path, err)
				} else {
					logger.Debugf("Added file %s to archive", path)
				}
			}

			return nil
		})

		// Set count to the number of files
		archiveFolder.Count = len(archiveFolder.Files)

		if err != nil {
			logger.Errorf("Error walking directory %s: %v", folderPath, err)
			continue
		}

		archiveSummary.Folders = append(archiveSummary.Folders, archiveFolder)
	}

	summaryFilename := fmt.Sprintf("archive_summary_%s.json", timestamp)
	archiveDir := AppConfig.DNSToolkit.Folders.Archive
	summaryPath := filepath.Join(archiveDir, summaryFilename)

	summaryJSON, err := json.MarshalIndent(archiveSummary, "", "  ")
	if err != nil {
		logger.Errorf("Failed to marshal archive summary: %v", err)
		return
	}

	if err := os.WriteFile(summaryPath, summaryJSON, 0644); err != nil {
		logger.Errorf("Failed to write archive summary to path %s: %v", summaryPath, err)
		return
	}

	logger.Infof("Archive process completed successfully; created archive %s", archivePath)
	logger.Infof("Summary file created at %s", summaryPath)
	logger.Infof("Total folders archived: %d", len(archiveSummary.Folders))
	logger.Infof("Total summary files archived: %d", len(archiveSummary.SummaryFiles))
}

// processSummaryFiles processes the summary files and adds them to the archive summary
func processSummaryFiles(
	logger *multilog.Logger,
	summaryDir string,
	archiveSummary *common.ArchiveSummary,
	_ *tar.Writer,
) {
	entries, err := os.ReadDir(summaryDir)
	if err != nil {
		logger.Errorf("Failed to read summary directory: %v", err)
		return
	}
	var summaryFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(name, "_summary.json") {
			summaryFiles = append(summaryFiles, filepath.Join(summaryDir, name))
		}
	}

	if len(summaryFiles) == 0 {
		logger.Warnf("No summary files found in %s", summaryDir)
		return
	}

	for _, path := range summaryFiles {
		content, err := os.ReadFile(path)
		if err != nil {
			logger.Errorf("Failed to read summary file: %v", err)
			continue
		}

		fileInfo, err := os.Stat(path)
		if err != nil {
			logger.Errorf("Failed to get file info: %v", err)
			continue
		}

		checksum := u.CalculateChecksumFromContent(
			content,
			AppConfig.DNSToolkit.FilesChecksum.Algorithm,
		)

		summaryType := u.DetermineSummaryTypeFromPath(path)

		count := getSummaryCount(content, summaryType)

		archiveSummaryFile := common.ArchiveSummaryFile{
			Name:        filepath.Base(path),
			Filepath:    path,
			SummaryType: summaryType,
			Count:       count,
			Checksum:    checksum,
			Size:        fileInfo.Size(),
			Valid:       true, // Assuming all summary files are valid
			Timestamp:   time.Now().Format(constants.TimestampFormat),
		}

		archiveSummary.SummaryFiles = append(archiveSummary.SummaryFiles, archiveSummaryFile)
	}
}

// getSummaryCount gets the count of items in the summary based on its type
func getSummaryCount(content []byte, summaryType string) int {
	switch summaryType {
	case constants.SummaryTypeDownload:
		var summary []common.DownloadSummary
		if err := json.Unmarshal(content, &summary); err == nil {
			return len(summary)
		}
	case constants.SummaryTypeProcessed:
		var summary []common.ProcessedSummary
		if err := json.Unmarshal(content, &summary); err == nil {
			return len(summary)
		}
	case constants.SummaryTypeConsolidated:
		var summary []common.ConsolidatedSummary
		if err := json.Unmarshal(content, &summary); err == nil {
			return len(summary)
		}
	case constants.SummaryTypeConsolidatedGroups:
		var summary []common.ConsolidatedGroupsSummary
		if err := json.Unmarshal(content, &summary); err == nil {
			return len(summary)
		}
	case constants.SummaryTypeOverlap:
		var summary []common.OverlapSummary
		if err := json.Unmarshal(content, &summary); err == nil {
			return len(summary)
		}
	case constants.SummaryTypeTop:
		var summary []common.TopSummary
		if err := json.Unmarshal(content, &summary); err == nil {
			return len(summary)
		}
	}

	return 0
}

// addFileToTar adds a file to the tar archive
func addFileToTar(
	logger *multilog.Logger,
	tarWriter *tar.Writer,
	sourcePath, targetPath string,
	info os.FileInfo,
) error {
	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}

	if targetPath == "" {
		targetPath = filepath.Base(sourcePath)
	}

	header.Name = targetPath

	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	if info.Mode().IsRegular() {
		file, err := os.Open(sourcePath)
		if err != nil {
			return err
		}
		defer func() {
			if err := file.Close(); err != nil {
				logger.Errorf("Failed to close file %s: %v\n", sourcePath, err)
			}
		}()

		_, err = io.Copy(tarWriter, file)
		if err != nil {
			return err
		}
	}

	return nil
}
