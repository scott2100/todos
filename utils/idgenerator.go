package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"todolist/utils/error"
)

func GenerateID() int {
	file, err := os.Open("todos.csv")
	error.HandleError(err)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	error.HandleError(err)

	fmt.Println("Records length: ", len(records))
	id := (len(records) - 1) + 1

	file.Close()
	return id
}
