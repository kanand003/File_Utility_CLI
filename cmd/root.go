package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-utility-cli",
	Short: "A CLI tool for various file operations",
	Long: `File Utility CLI is a command-line tool that provides various file operations
such as file manipulation, directory operations, and content processing.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("File Utility CLI starting...")
		logFile, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Error opening log file: %v\n", err)
			return
		}
		log.SetOutput(logFile)
		log.Println("CLI execution started")
	},
}

func Execute() {
	fmt.Println("Execute function called")
	log.Println("Execute function called")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing root command:", err)
		log.Printf("Error executing root command: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}
