package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"todolist/todo"
	"todolist/utils/error"
	"todolist/utils/file"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete todo",
	Long:  `Removes todo from the todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		var updatedTodos []todo.Todo

		if len(args) <= 0 {
			println("No argument provided. You must specify the ID of the task to delete.")
			return
		}

		rowIDToRemove, err := strconv.Atoi(args[0])
		error.HandleError(err)

		todos := file.ReadFile()

		for _, currentTodo := range todos {
			if currentTodo.ID != rowIDToRemove {
				updatedTodos = append(updatedTodos, todo.Todo{ID: currentTodo.ID, Description: currentTodo.Description,
					Created: currentTodo.Created, Completed: currentTodo.Completed})
			}
		}

		file.UpdateFile(updatedTodos)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().BoolP("completed", "c", false, "Remove all completed tasks.")
	deleteCmd.PersistentFlags().BoolP("uncompleted", "u", false, "Remove all uncompleted tasks.")
	deleteCmd.PersistentFlags().BoolP("all", "a", false, "Delete all tasks.")
}
