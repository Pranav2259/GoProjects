package todo

import "fmt"

func Help() {
	cyan := "\033[36m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	// ASCII art in cyan
	fmt.Println(cyan + `___________   ________              _____
\__    ___/___\______ \   ____     /  _  \ ______ ______
  |    | /  _ \|    |  \ /  _ \   /  /_\  \\____ \\____ \
  |    |(  <_> )    \  (  <_> ) /    |    \  |_> >  |_> >
  |____| \____/_______  /\____/  \____|__  /   __/|   __/
                      \/                 \/|__|   |__|
` + reset)

	// App title in green
	fmt.Println(green + "todolist - CLI based task tracking App" + reset)
	fmt.Println()

	// Usage in yellow
	fmt.Println(yellow + "Usage:" + reset)
	fmt.Println("  todolist add \"task description\"       # Add a new task")
	fmt.Println("  todolist list                           # List pending tasks")
	fmt.Println("  todolist list-all                       # List all tasks")
	fmt.Println("  todolist complete <id>                  # Mark task as completed")
	fmt.Println("  todolist delete <id>                    # Delete a task")
	fmt.Println("  todolist help                           # Show this help message")
}
