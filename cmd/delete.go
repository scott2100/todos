package cmd

import (
	"github.com/spf13/cobra"
	"todolist/todo"
	"todolist/utils/file"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete todo",
	Long:  `Removes todo from the todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		var todos []todo.Todo
		var rowID string
		if len(args) > 0 {
			rowID = args[0]
		} else {
			println("No argument provided. You must specify the ID of the task to delete.")
			return
		}

		todos = file.ReadFile(todos, rowID)
		file.UpdateFile(todos)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().BoolP("completed", "c", false, "Remove all completed tasks.")
	deleteCmd.PersistentFlags().BoolP("uncompleted", "u", false, "Remove all uncompleted tasks.")
	deleteCmd.PersistentFlags().BoolP("all", "a", false, "Delete all tasks.")
}
