package cmd

import (
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "File Operations",
	Long:  "Perform various file operations like copy,move and delete",
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
