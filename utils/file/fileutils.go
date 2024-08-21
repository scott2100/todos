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

func UpdateFile(todos []todo.Todo) {
	file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	error.CheckError(err)
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

func ReadFile(todos []todo.Todo, rowID string) []todo.Todo {
	file, err := os.Open("todos.csv")
	error.CheckError(err)

	defer file.Close()

	r, err := csv.NewReader(file).ReadAll()
	error.CheckError(err)

	for _, row := range r[1:] {
		id, err := strconv.Atoi(row[0])
		error.CheckError(err)
		createdTime, err := time.Parse(time.RFC3339, row[2])
		error.CheckError(err)
		isCompleted, err := strconv.ParseBool(row[3])
		error.CheckError(err)

		if row[0] != rowID {
			todos = append(todos, todo.Todo{ID: id, Description: row[1], Created: createdTime, IsComplete: isCompleted})
		}
	}
	return todos
}
