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
}

func check(e error) {
	if e != nil {
		log.Fatal("error occurred reading csv file: ", e)
	}
}
