package todo

import (
	"fmt"
	"strconv"

	"github.com/Pranav2259/GoProjects/todolist/storage"
)

func Complete(args []string) {
	if len(args) == 0 {
		fmt.Println("❌ Task ID required")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("❌ Invalid task ID")
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("❌ Failed to load tasks:", err)
		return
	}

	updated := false
	for _, task := range tasks {
		if task.ID == id {
			task.Status = storage.Completed
			if err := storage.UpdateTask(task); err != nil {
				fmt.Println("❌ Failed to update task:", err)
			} else {
				fmt.Println("✅ Task marked as completed")
			}
			updated = true
			break
		}
	}

	if !updated {
		fmt.Println("❌ Task not found")
	}
}
