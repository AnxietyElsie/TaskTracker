package service

import (
	"TaskTracker/internal/config"
	"TaskTracker/internal/storage"
	"fmt"
	"time"
)

func UpdateTask(args []string, id int) error {
	tasks := storage.LoadTasks()

	if len(args) < 4 {
		return fmt.Errorf("To update your task's status use format: 'task-cli update [id] \"Your changed task\"'")
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = args[3]
			err := storage.SaveTasks(tasks)
			if err != nil {
				return fmt.Errorf("Error saving task list: %v", err)
			}
			tasks[i].UpdatedAt = time.Now().Format(config.Format)
			fmt.Println("Your task successfully updated!")
			return nil
		}
	}
	return fmt.Errorf("ID not found")
}
