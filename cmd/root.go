package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todolist",
	Short: "A todo list command line application",
	Long:  `Todo list application that runs via command line. You can add, delete and edit todo list items.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func check(e error) {
	if e != nil {
		log.Fatal("Error occurred reading csv file: ", e)
	}
}
