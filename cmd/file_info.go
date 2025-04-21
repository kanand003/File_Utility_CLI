package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var fileInfoCmd = &cobra.Command{
	Use:   "info [file]",
	Short: "Display detailed information about a file",
	Long:  "Display detailed information about a file including size, permissions, creation time, and more",
	Args:  cobra.ExactArgs(1),
	RunE:  runFileInfo,
}

func init() {
	fileCmd.AddCommand(fileInfoCmd)
}

func runFileInfo(cmd *cobra.Command, args []string) error {
	target := args[0]

	fmt.Printf("Gathering information for: %s\n", target)

	// Get file info
	fileInfo, err := os.Stat(target)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File not found: %s\n", target)
			return fmt.Errorf("file not found: %s", target)
		}
		fmt.Printf("Error accessing file: %v\n", err)
		return fmt.Errorf("error accessing file: %w", err)
	}

	// Get absolute path
	absPath, err := filepath.Abs(target)
	if err != nil {
		fmt.Printf("Warning: Could not determine absolute path: %v\n", err)
		absPath = target
	}

	// Display file information
	fmt.Println("\nFile Information:")
	fmt.Printf("%-20s: %s\n", "Name", fileInfo.Name())
	fmt.Printf("%-20s: %s\n", "Absolute Path", absPath)
	fmt.Printf("%-20s: %s\n", "Type", getFileType(fileInfo))
	fmt.Printf("%-20s: %d bytes (%.2f KB)\n", "Size", fileInfo.Size(), float64(fileInfo.Size())/1024)
	fmt.Printf("%-20s: %s\n", "Permissions", fileInfo.Mode().String())
	fmt.Printf("%-20s: %s\n", "Last Modified", fileInfo.ModTime().Format(time.RFC1123))

	// If it's a symlink, show the target
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		linkTarget, err := os.Readlink(target)
		if err != nil {
			fmt.Printf("%-20s: Error reading link target: %v\n", "Link Target", err)
		} else {
			fmt.Printf("%-20s: %s\n", "Link Target", linkTarget)
		}
	}

	// Additional file system information if available
	if sys := fileInfo.Sys(); sys != nil {
		fmt.Println("\nAdditional system-specific information may be available but requires type assertion")
	}

	return nil
}

func getFileType(info os.FileInfo) string {
	if info.IsDir() {
		return "Directory"
	}
	mode := info.Mode()
	if mode&os.ModeSymlink != 0 {
		return "Symbolic Link"
	}
	if mode&os.ModeDevice != 0 {
		return "Device"
	}
	if mode&os.ModeNamedPipe != 0 {
		return "Named Pipe"
	}
	if mode&os.ModeSocket != 0 {
		return "Socket"
	}
	if mode&os.ModeCharDevice != 0 {
		return "Character Device"
	}
	return "Regular File"
}
