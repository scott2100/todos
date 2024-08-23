package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"todolist/todo"
	"todolist/utils/file"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark todo as completed",
	Long:  `Mark todo as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		var updatedTodos []todo.Todo
		rowID, err := strconv.Atoi(args[0])
		error.HandleError(err)

		if len(args) <= 0 {
			println("No argument provided. You must specify the ID of the task to mark as completed.")
			return
		}

		todos := file.ReadFile()

		for _, currentTodo := range todos {
			if currentTodo.ID == rowID {
				currentTodo.IsComplete = true
			}
			updatedTodos = append(updatedTodos, todo.Todo{ID: currentTodo.ID, Description: currentTodo.Description,
				Created: currentTodo.Created, IsComplete: currentTodo.IsComplete})
		}

		file.UpdateFile(updatedTodos)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
