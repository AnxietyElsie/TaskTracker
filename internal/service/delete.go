package service

import (
	"TaskTracker/internal/storage"
	"fmt"
)

func DeleteTask(id int) error {
	tasks := storage.LoadTasks()
	for i, t := range tasks {
		if t.ID == id {
			newTaskList := append(tasks[:i], tasks[i+1:]...)
			err := storage.SaveTasks(newTaskList)
			if err != nil {
				return fmt.Errorf("Error saving task list: %v", err)
			}
			fmt.Printf("Task Deleted! [ID: %d]\n", t.ID)
			return nil
		}
	}
	return fmt.Errorf("ID not found")
}
