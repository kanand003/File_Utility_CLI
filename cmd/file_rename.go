package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var fileRenameCmd = &cobra.Command{
	Use:   "rename [old_name] [new_name]",
	Short: "Rename a file",
	Long:  "Rename a file from old name to new name",
	Args:  cobra.ExactArgs(2),
	RunE:  runFileRename,
}

func init() {
	fileCmd.AddCommand(fileRenameCmd)
}

func runFileRename(cmd *cobra.Command, args []string) error {
	oldName := args[0]
	newName := args[1]

	fmt.Printf("Renaming file from %s to %s\n", oldName, newName)

	// Check if source exists
	fileInfo, err := os.Stat(oldName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File not found: %s\n", oldName)
			return fmt.Errorf("file not found: %s", oldName)
		}
		fmt.Printf("Error accessing file: %v\n", err)
		return fmt.Errorf("error accessing file: %w", err)
	}

	// Check if it's a directory
	if fileInfo.IsDir() {
		return fmt.Errorf("source is a directory, use 'dir rename' instead")
	}

	// Check if destination already exists
	if _, err := os.Stat(newName); err == nil {
		fmt.Printf("Warning: Destination file %s already exists and will be overwritten\n", newName)
	}

	// Create destination directory if needed
	destDir := filepath.Dir(newName)
	if destDir != "." && destDir != filepath.Dir(oldName) {
		fmt.Printf("Creating directory: %s\n", destDir)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return fmt.Errorf("error creating destination directory: %w", err)
		}
	}

	// Rename the file
	if err := os.Rename(oldName, newName); err != nil {
		fmt.Printf("Error renaming file: %v\n", err)
		return fmt.Errorf("error renaming file: %w", err)
	}

	fmt.Printf("Successfully renamed %s to %s\n", oldName, newName)
	return nil
}
