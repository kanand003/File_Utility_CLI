package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var fileCopyCmd = &cobra.Command{
	Use:   "copy [source] [destination]",
	Short: "Copy a file from source to destination",
	Args:  cobra.ExactArgs(2),
	RunE:  runFileCopy,
}

func init() {
	fileCmd.AddCommand(fileCopyCmd)
}

func runFileCopy(cmd *cobra.Command, args []string) error {
	source := args[0]
	destination := args[1]

	fmt.Printf("Copying from %s to %s\n", source, destination)

	// Check if source exists
	sourceInfo, err := os.Stat(source)
	if err != nil {
		fmt.Printf("Error accessing source file: %v\n", err)
		return fmt.Errorf("error accessing source file: %w", err)
	}
	fmt.Printf("Source file size: %d bytes\n", sourceInfo.Size())

	if sourceInfo.IsDir() {
		return fmt.Errorf("source is a directory, use 'dir copy' instead")
	}

	// Create destination directory
	destDir := filepath.Dir(destination)
	fmt.Printf("Creating directory: %s\n", destDir)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return fmt.Errorf("error creating destination directory: %w", err)
	}

	// Open source file
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Printf("Error opening source: %v\n", err)
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer sourceFile.Close()

	// Create destination file
	destFile, err := os.Create(destination)
	if err != nil {
		fmt.Printf("Error creating destination: %v\n", err)
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer destFile.Close()

	// Copy the contents
	fmt.Println("Copying file contents...")
	bytesWritten, err := io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Printf("Error during copy: %v\n", err)
		return fmt.Errorf("error copying file: %w", err)
	}
	fmt.Printf("Copied %d bytes\n", bytesWritten)

	// Copy file permissions
	if err := os.Chmod(destination, sourceInfo.Mode()); err != nil {
		fmt.Printf("Error setting permissions: %v\n", err)
		return fmt.Errorf("error setting file permissions: %w", err)
	}

	fmt.Printf("Successfully copied %s to %s\n", source, destination)
	return nil
}
