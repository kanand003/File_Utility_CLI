package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of File Utility CLI",
	Long:  `All software has versions. This is File Utility CLI's version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("File Utility CLI v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
