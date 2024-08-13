package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func GenerateID() int {
	file, err := os.Open("todos.csv")
	check(err)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	check(err)

	fmt.Println("Records length: ", len(records))
	id := (len(records) - 1) + 1

	return id
}

func check(e error) {
	if e != nil {
		log.Fatal("error occurred reading csv file: ", e)
	}
}
