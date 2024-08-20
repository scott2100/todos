package main

import (
	"encoding/csv"
	"log"
	"os"
	"todolist/cmd"
)

func main() {
	file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	check(err)
	defer file.Close()

	fileInfo, err := file.Stat()

	w := csv.NewWriter(file)

	headers := []string{"ID", "DESCRIPTION", "CREATED", "COMPLETED"}
	if fileInfo.Size() == 0 {
		err = w.Write(headers)
		w.Flush()
	}
	cmd.Execute()
}

func check(e error) {
	if e != nil {
		log.Fatal("error occurred reading csv file: ", e)
	}
}
