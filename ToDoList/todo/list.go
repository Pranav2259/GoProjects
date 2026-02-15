package todo

import (
	"fmt"

	"github.com/Pranav2259/GoProjects/todolist/output"
	"github.com/Pranav2259/GoProjects/todolist/storage"
)

func List(args []string) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("❌ Failed to load tasks:", err)
		return
	}

	var pending []storage.Task
	for _, t := range tasks {
		if t.Status == storage.Pending {
			pending = append(pending, t)
		}
	}

	output.PrintTasks(pending)
}

func ListAll(args []string) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("❌ Failed to load tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	output.PrintTasks(tasks)
}
