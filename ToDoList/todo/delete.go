package todo

import (
	"fmt"
	"strconv"

	"github.com/Pranav2259/GoProjects/todolist/storage"
)

func Delete(args []string) {
	if len(args) == 0 {
		fmt.Println("❌ Task ID required")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("❌ Invalid task ID")
		return
	}

	if err := storage.DeleteTask(id); err != nil {
		fmt.Println("❌ Failed to delete task:", err)
		return
	}

	fmt.Println("🗑 Task deleted")
}
