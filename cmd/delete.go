package cmd

import (
	"github.com/spf13/cobra"
	"todolist/utils/database"
	"todolist/utils/error"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete todo",
	Long:  `Removes todo from the todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			println("No argument provided. You must specify the ID of the task to delete.")
			return
		}

		db := database.OpenDBConnection()
		defer db.Close()

		deleteSql := `DELETE FROM todos WHERE id = ?`
		statement, err := db.Prepare(deleteSql)
		error.HandleError(err)
		_, err = statement.Exec(args[0])
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().BoolP("completed", "c", false, "Remove all completed tasks.")
	deleteCmd.PersistentFlags().BoolP("uncompleted", "u", false, "Remove all uncompleted tasks.")
	deleteCmd.PersistentFlags().BoolP("all", "a", false, "Delete all tasks.")
}

func deleteFromCSV() {
	/*
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

	*/
}
