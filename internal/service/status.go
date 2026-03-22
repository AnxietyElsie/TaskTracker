package service

import (
	"TaskTracker/internal/config"
	"TaskTracker/internal/storage"
	"fmt"
	"time"
)

func UpdateStatus(args []string, id int) error {
	tasks := storage.LoadTasks()

	if len(args) < 2 {
		return fmt.Errorf("To update your task's status use format: 'task-cli <mark> [id]'")
	}

	for i, t := range tasks {
		if t.ID == id {
			switch args[2] {
			case "done":
				tasks[i].Status = "done"
			case "in-progress":
				tasks[i].Status = "in-progress"
			default:
				return fmt.Errorf("Unknown status")
			}
			tasks[i].UpdatedAt = time.Now().Format(config.Format)
			err := storage.SaveTasks(tasks)
			if err != nil {
				return fmt.Errorf("Error saving task list: %v", err)
			}
			fmt.Println("Task status updated!")
			return nil
		}
	}
	return fmt.Errorf("ID not found")
}
