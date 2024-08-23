package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func GenerateID() int {
	file, err := os.Open("todos.csv")
	check(err)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	check(err)

	id := (len(records) - 1) + 1

	file.Close()
	return id
}

func check(e error) {
	if e != nil {
		log.Fatal("Error occurred reading csv file: ", e)
	}
}
