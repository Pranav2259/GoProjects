package todo

import (
	"fmt"
	"time"

	"github.com/Pranav2259/GoProjects/todolist/storage"
)

func Add(args []string) {
	if len(args) == 0 {
		fmt.Println("❌ Task description is required")
		fmt.Println(`Usage: tasks add "task description"`)
		return
	}

	taskText := args[0]

	id, err := storage.NextID()
	if err != nil {
		fmt.Println("❌ Error generating ID:", err)
		return
	}

	task := storage.Task{
		ID:        id,
		Task:      taskText,
		Status:    storage.Pending,
		CreatedAt: time.Now().UTC(),
	}

	if err := storage.SaveTask(task); err != nil {
		fmt.Println("❌ Failed to save task:", err)
		return
	}

	fmt.Println("✅ Task added successfully")
	fmt.Println("ID:", task.ID)
	fmt.Println("Task:", task.Task)
	fmt.Println("Status:", task.Status)
	fmt.Println("Created:", task.CreatedAt.Format("02 Jan 2006 15:04"))
}
