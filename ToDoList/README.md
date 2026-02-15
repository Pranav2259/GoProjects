# ToDoList CLI

A simple CLI-based To-Do List application written in Go.  
This project is built to learn Go CLI development step by step using
standard libraries first.

## Goals
- Learn Go project structure
- Understand CLI argument parsing
- Practice clean Git workflow
- Build an installable CLI using `go install`

## Planned Features
- Add a new todo item
- List all todo items
- Mark a todo as done
- Persist todos locally (CSV file)
- Basic help and usage output
- Color-coded output for pending/completed tasks

## Installation

1. **Clone the repository**
```bash
git clone https://github.com/Pranav2259/GoProjects.git
cd GoProjects/ToDoList
```
2. **Install the CLI using Go**
```bash
go install
```

## CLI USAGE

- todolist                   Show help / usage
- todolist add "title"        Create a new task
- todolist list               List pending tasks
- todolist list-all           List all tasks
- todolist complete <id>      Mark a task as completed
- todolist delete <id>        Delete a task
- todolist help               Show this help message

## EXAMPLE 

- todolist                   Show help / usage
- todolist add "title"        Create a new task
- todolist list               List pending tasks
- todolist list-all           List all tasks
- todolist complete <id>      Mark a task as completed
- todolist delete <id>        Delete a task
- todolist help               Show this help message
