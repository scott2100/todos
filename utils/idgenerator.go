package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func GenerateID() int {
	file, err := os.Open("todos.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Records length: ", len(records))
	id := (len(records) - 1) + 1

	file.Close()
	return id
}
