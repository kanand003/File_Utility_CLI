package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var fileMoveCmd = &cobra.Command{
	Use:   "move [source] [destination]",
	Short: "Move a file from source to destination",
	Long:  "Move a file from source path to destination path. This operation removes the source file after successful transfer.",
	Args:  cobra.ExactArgs(2),
	RunE:  runFileMove,
}

func init() {
	fileCmd.AddCommand(fileMoveCmd)
}

func runFileMove(cmd *cobra.Command, args []string) error {
	source := args[0]
	destination := args[1]

	fmt.Printf("Moving from %s to %s\n", source, destination)

	// Check if source exists
	sourceInfo, err := os.Stat(source)
	if err != nil {
		fmt.Printf("Error accessing source file: %v\n", err)
		return fmt.Errorf("error accessing source file: %w", err)
	}
	fmt.Printf("Source file size: %d bytes\n", sourceInfo.Size())

	if sourceInfo.IsDir() {
		return fmt.Errorf("source is a directory, use 'dir move' instead")
	}

	// Create destination directory if it doesn't exist
	destDir := filepath.Dir(destination)
	fmt.Printf("Creating directory: %s\n", destDir)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return fmt.Errorf("error creating destination directory: %w", err)
	}

	// Check if destination already exists
	if _, err := os.Stat(destination); err == nil {
		fmt.Printf("Warning: Destination file %s already exists and will be overwritten\n", destination)
	}

	// First try to use os.Rename which is more efficient for same filesystem moves
	err = os.Rename(source, destination)
	if err == nil {
		fmt.Printf("Successfully moved %s to %s\n", source, destination)
		return nil
	}

	// If rename failed (e.g., cross-device link), fall back to copy and delete
	fmt.Printf("Direct move failed: %v\n", err)
	fmt.Println("Falling back to copy and delete method...")

	// Copy the file
	if err := copyFile(source, destination); err != nil {
		return fmt.Errorf("error copying file during move operation: %w", err)
	}

	// Remove the source file
	fmt.Printf("Removing source file: %s\n", source)
	if err := os.Remove(source); err != nil {
		fmt.Printf("Warning: Could not remove source file: %v\n", err)
		return fmt.Errorf("moved file successfully but failed to remove source file: %w", err)
	}

	fmt.Printf("Successfully moved %s to %s\n", source, destination)
	return nil
}

// Helper function to copy a file
func copyFile(src, dst string) error {
	// Open source file
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("Error opening source: %v\n", err)
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer sourceFile.Close()

	// Get source file info for permissions
	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		fmt.Printf("Error getting source file info: %v\n", err)
		return fmt.Errorf("error getting source file info: %w", err)
	}

	// Create destination file
	destFile, err := os.Create(dst)
	if err != nil {
		fmt.Printf("Error creating destination: %v\n", err)
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer destFile.Close()

	// Copy the contents
	fmt.Println("Copying file contents...")
	bytesWritten, err := destFile.ReadFrom(sourceFile)
	if err != nil {
		fmt.Printf("Error during copy: %v\n", err)
		return fmt.Errorf("error copying file: %w", err)
	}
	fmt.Printf("Copied %d bytes\n", bytesWritten)

	// Copy file permissions
	if err := os.Chmod(dst, sourceInfo.Mode()); err != nil {
		fmt.Printf("Error setting permissions: %v\n", err)
		return fmt.Errorf("error setting file permissions: %w", err)
	}

	return nil
}
