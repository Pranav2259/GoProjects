package main

import (
	"os"
	"strings"

	todo "github.com/Pranav2259/GoProjects/todolist/todo"
)

func main() {
	if len(os.Args) < 2 {
		todo.Help()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch strings.ToLower(command) {
	case "add":
		todo.Add(args)
	case "delete":
		todo.Delete(args)
	case "complete":
		todo.Complete(args)
	case "list":
		todo.List(args)
	case "list-all":
		todo.ListAll(args)
	case "help":
		todo.Help()
	default:
		todo.Help()
	}
}
