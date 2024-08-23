package main

import (
	"todolist/cmd"
	"todolist/utils/file"
)

func main() {
	file.WriteHeaders()
	cmd.Execute()
}
