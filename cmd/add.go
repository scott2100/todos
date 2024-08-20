package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
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
		//fmt.Println("add called")

		file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_APPEND, 0644)
		check(err)
		defer file.Close()

		w := csv.NewWriter(file)
		todo := todo.Todo{ID: utils.GenerateID(), Description: strings.Join(args, ""), Created: time.Now(), IsComplete: false}
		fmt.Println("Todo: ", todo)
		w.Write(todo.Slice())
		w.Flush()

		fmt.Println("Data written to the CSV file successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
