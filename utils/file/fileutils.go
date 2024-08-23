package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	"todolist/todo"
	"todolist/utils/error"
)

const Header = "ID\tTasks\tCreated\tCompleted"

func UpdateFile(todos []todo.Todo) {
	file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	error.HandleError(err)
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

func ReadFile() []todo.Todo {
	var todos []todo.Todo

	file, err := os.Open("todos.csv")
	error.HandleError(err)

	defer file.Close()

	r, err := csv.NewReader(file).ReadAll()
	error.HandleError(err)

	for _, row := range r[1:] {
		id, err := strconv.Atoi(row[0])
		error.HandleError(err)
		createdTime, err := time.Parse(time.RFC3339, row[2])
		error.HandleError(err)
		isCompleted, err := strconv.ParseBool(row[3])
		error.HandleError(err)

		todoToAppend := todo.Todo{ID: id, Description: row[1], Created: createdTime, IsComplete: isCompleted}

		todos = append(todos, todoToAppend)
	}
	return todos
}

func WriteHeaders() {
	file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	error.HandleError(err)
	defer file.Close()

	fileInfo, err := file.Stat()
	error.HandleError(err)

	w := csv.NewWriter(file)

	headers := []string{"ID", "Tasks", "Created", "Completed"}
	if fileInfo.Size() == 0 {
		err = w.Write(headers)
		error.HandleError(err)
		w.Flush()
	}
}
