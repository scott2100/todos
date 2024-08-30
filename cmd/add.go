package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"time"
	"todolist/todo"
	"todolist/utils"
	"todolist/utils/database"
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
	newTodo := todo.Todo{ID: utils.GenerateID(), Description: strings.Join(args, ""), Created: time.Now(), Completed: time.Time{}}

	db := database.OpenDBConnection()
	defer db.Close()

	insertSql := `INSERT INTO todos(description, created, completed) values (?, ?, ?)`
	preparedStatement, err := db.Prepare(insertSql)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = preparedStatement.Exec(newTodo.Description, time.Now(), time.Time{})
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
