package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
	"todolist/todo"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete todo",
	Long:  `Removes todo from the todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		var todos []todo.Todo
		var rowToDelete string
		if len(args) > 0 {
			rowToDelete = args[0]
		} else {
			println("No argument provided. You must specify the ID of the task to delete.")
			return
		}

		todos = readFile(todos, rowToDelete)
		updateFile(todos)
	},
}

func updateFile(todos []todo.Todo) {
	file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	check(err)
	defer file.Close()

	fileInfo, err := file.Stat()
	w := csv.NewWriter(file)
	headers := []string{"ID", "DESCRIPTION", "CREATED", "COMPLETED"}
	if fileInfo.Size() == 0 {
		err = w.Write(headers)
		w.Flush()
	}
	for _, todo := range todos {
		w.Write(todo.Slice())
	}
	w.Flush()

	fmt.Println("Data written to the CSV file successfully!")
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().BoolP("completed", "c", false, "Remove all completed tasks.")
	deleteCmd.PersistentFlags().BoolP("uncompleted", "u", false, "Remove all uncompleted tasks.")
	deleteCmd.PersistentFlags().BoolP("all", "a", false, "Delete all tasks.")
}

func readFile(todos []todo.Todo, rowToDelete string) []todo.Todo {
	println("In read file")
	file, err := os.Open("todos.csv")
	check(err)

	defer file.Close()

	r, err := csv.NewReader(file).ReadAll()
	check(err)

	for _, row := range r[1:] {
		id, err := strconv.Atoi(row[0])
		check(err)
		time, err := time.Parse(time.RFC3339, row[2])
		check(err)
		isCompleted, err := strconv.ParseBool(row[3])
		check(err)
		println("ID -> ", row[0])
		if row[0] != rowToDelete {
			println("In append todos")
			todos = append(todos, todo.Todo{ID: id, Description: row[1], Created: time, IsComplete: isCompleted})
		}
	}
	return todos
}
