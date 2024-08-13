/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"time"
	"todolist/todo"
	"todolist/utils"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the todo list.",
	Long: `Adds a new task to the todolist.

For example:

todo add Buy Milk.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		file, err := os.OpenFile("todos.csv", os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		/*
			// this defines the header value and data values for the new csv file
			headers := []string{"ID", "TODO", "CREATED"}
			w.Write(headers)
		*/

		check(err)
		defer file.Close()

		w := csv.NewWriter(file)

		todo := todo.Todo{ID: utils.GenerateID(), Text: strings.Join(args, ""), Created: time.Now()}
		record := todo.Slice()
		w.Write(record)

		// Write any buffered data to the underlying writer (standard output).
		w.Flush()

		err = w.Error()
		check(err)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func check(e error) {
	if e != nil {
		log.Fatal("error occurred reading csv file: ", e)
	}
}
