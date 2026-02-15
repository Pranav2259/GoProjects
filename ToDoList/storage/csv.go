package storage

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var fileName string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot determine user home directory: " + err.Error())
	}

	// Create a hidden folder .todolist in the home directory
	folder := filepath.Join(homeDir, ".todolist")
	err = os.MkdirAll(folder, 0755)
	if err != nil {
		panic("Cannot create .todolist folder: " + err.Error())
	}

	// Full path to CSV file
	fileName = filepath.Join(folder, "tasks.csv")
}

// ensureFile checks if the CSV exists, and creates it with headers if missing
func ensureFile() error {
	if _, err := os.Stat(fileName); err == nil {
		// File exists
		return nil
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	return writer.Write([]string{"id", "task", "status", "created_at"})
}

// GetFileName returns the full CSV path
func GetFileName() string {
	return fileName
}

func LoadTasks() ([]Task, error) {
	if err := ensureFile(); err != nil {
		return nil, err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		id, _ := strconv.Atoi(row[0])
		createdAt, _ := time.Parse(time.RFC3339, row[3])

		status := Pending
		if row[2] == "completed" {
			status = Completed
		}

		tasks = append(tasks, Task{
			ID:        id,
			Task:      row[1],
			Status:    status,
			CreatedAt: createdAt,
		})
	}

	return tasks, nil
}
func NextID() (int, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return 0, err
	}

	if len(tasks) == 0 {
		return 1, nil
	}

	return tasks[len(tasks)-1].ID + 1, nil
}
func SaveTask(task Task) error {
	if err := ensureFile(); err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write([]string{
		strconv.Itoa(task.ID),
		task.Task,
		task.Status.String(),
		task.CreatedAt.Format(time.RFC3339),
	})
}
func UpdateTask(updated Task) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"id", "task", "status", "created_at"})

	for _, task := range tasks {
		if task.ID == updated.ID {
			task = updated
		}

		writer.Write([]string{
			strconv.Itoa(task.ID),
			task.Task,
			task.Status.String(),
			task.CreatedAt.Format(time.RFC3339),
		})
	}

	return nil
}
func DeleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"id", "task", "status", "created_at"})

	for _, task := range tasks {
		if task.ID == id {
			continue
		}

		writer.Write([]string{
			strconv.Itoa(task.ID),
			task.Task,
			task.Status.String(),
			task.CreatedAt.Format(time.RFC3339),
		})
	}

	return nil
}
