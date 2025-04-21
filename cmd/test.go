package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test if the CLI is working",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test command executed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
