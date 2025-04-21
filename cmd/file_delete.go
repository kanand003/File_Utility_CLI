package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var fileDeleteCmd = &cobra.Command{
	Use:   "delete [file]",
	Short: "Delete a file",
	Long:  "Delete a file from the filesystem",
	Args:  cobra.ExactArgs(1),
	RunE:  runFileDelete,
}

func init() {
	fileCmd.AddCommand(fileDeleteCmd)
}

func runFileDelete(cmd *cobra.Command, args []string) error {
	target := args[0]

	fmt.Printf("Attempting to delete file: %s\n", target)

	// Check if file exists
	fileInfo, err := os.Stat(target)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File not found: %s\n", target)
			return fmt.Errorf("file not found: %s", target)
		}
		fmt.Printf("Error accessing file: %v\n", err)
		return fmt.Errorf("error accessing file: %w", err)
	}

	// Check if it's a directory
	if fileInfo.IsDir() {
		return fmt.Errorf("target is a directory, use 'dir delete' instead")
	}

	// Delete the file
	if err := os.Remove(target); err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return fmt.Errorf("error deleting file: %w", err)
	}

	fmt.Printf("Successfully deleted file: %s\n", target)
	return nil
}
