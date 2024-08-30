package cmd

import (
	"github.com/spf13/cobra"
	"time"
	"todolist/utils/database"
	"todolist/utils/error"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark todo as completed",
	Long:  `Mark todo as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			println("No argument provided. You must specify the ID of the task to delete.")
			return
		}

		db := database.OpenDBConnection()
		defer db.Close()

		markCompleteSql := `UPDATE todos SET completed = ? where id = ?`
		statement, err := db.Prepare(markCompleteSql)
		error.HandleError(err)
		//completedTime := time.Now().Format(time.RFC822)
		completedTime := time.Now()
		_, err = statement.Exec(completedTime, args[0])
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func markCompleteInCSV() {
	/*
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
				currentTodo.Completed = time.Now()
			}
			updatedTodos = append(updatedTodos, todo.Todo{ID: currentTodo.ID, Description: currentTodo.Description,
				Created: currentTodo.Created, Completed: currentTodo.Completed})
		}

		file.UpdateFile(updatedTodos)

	*/
}
