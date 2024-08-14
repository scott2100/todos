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
		headers := []string{"ID", "TODO", "CREATED"}

		file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		check(err)
		defer file.Close()

		fileInfo, err := file.Stat()
		check(err)

		w := csv.NewWriter(file)

		if fileInfo.Size() == 0 {
			err = w.Write(headers)
			check(err)
			w.Flush()
			fmt.Println("CSV headers written to the file.")
		} else {
			fmt.Println("File already contains data, skipping header writing.")
		}
		todo := todo.Todo{ID: utils.GenerateID(), Text: strings.Join(args, ""), Created: time.Now()}
		todoSlice := todo.Slice()
		err = w.Write(todoSlice)
		check(err)
		w.Flush()

		fmt.Println("Data written to the CSV file successfully!")
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
