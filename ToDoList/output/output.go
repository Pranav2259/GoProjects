package output

import (
	"fmt"
	"strings"

	"github.com/Pranav2259/GoProjects/todolist/storage"
)

// PrintTasks prints a list of tasks in a table
func PrintTasks(tasks []storage.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	// ANSI colors
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	// Table header
	fmt.Printf("%-5s | %-50s | %-10s | %-20s\n", "ID", "Task", "Status", "Created At")
	fmt.Println(strings.Repeat("-", 90))

	// Rows
	for _, t := range tasks {
		statusStr := ""
		switch t.Status {
		case storage.Pending:
			statusStr = yellow + "Pending" + reset
		case storage.Completed:
			statusStr = green + "Completed" + reset
		}

		fmt.Printf("%-5d | %-50s | %-10s | %-20s\n",
			t.ID,
			t.Task,
			statusStr,
			t.CreatedAt.Format("02 Jan 2006 15:04"))
	}
}
