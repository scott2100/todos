package main

import (
	"todolist/cmd"
	"todolist/utils/database"
)

func main() {
	database.CreateDatabase()
	cmd.Execute()
}
