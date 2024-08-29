package cmd

import (
	"encoding/csv"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"time"
	"todolist/todo"
	"todolist/utils"
	"todolist/utils/database"
	"todolist/utils/error"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the todo list.",
	Long: `Adds a new task to the todolist.

For example:

todo add Buy Milk.`,
	Run: func(cmd *cobra.Command, args []string) {
		addToDB(args)
	},
}

func addToDB(args []string) {
	db := database.OpenDBConnection()
	defer db.Close()
	insertSql := `INSERT INTO todos(description, created, completed) values (?, ?, ?)`
	todo := todo.Todo{ID: utils.GenerateID(), Description: strings.Join(args, ""), Created: time.Now(), Completed: time.Time{}}
	preparedStatement, err := db.Prepare(insertSql)
	if err != nil {
		log.Fatal(err)
	}
	_, err = preparedStatement.Exec(todo.Description, time.Now(), time.Time{})
	if err != nil {
		log.Fatal(err)
	}
}

// consider removing or make using csv optional
func addToCSV(args []string) {
	file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_APPEND, 0644)
	error.HandleError(err)
	defer file.Close()

	w := csv.NewWriter(file)
	todo := todo.Todo{ID: utils.GenerateID(), Description: strings.Join(args, ""), Created: time.Now(), Completed: time.Time{}}

	w.Write(todo.Slice())
	w.Flush()
}

func init() {
	rootCmd.AddCommand(addCmd)
}
