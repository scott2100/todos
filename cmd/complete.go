package cmd

import (
	"github.com/spf13/cobra"
	"todolist/todo"
	"todolist/utils/file"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark todo as completed",
	Long:  `Mark todo as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		var todos []todo.Todo
		var rowID string
		if len(args) > 0 {
			rowID = args[0]
		} else {
			println("No argument provided. You must specify the ID of the task to mark as completed.")
			return
		}

		todos = file.ReadFile(todos, rowID)
		file.UpdateFile(todos)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
