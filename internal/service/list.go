package service

import (
	"TaskTracker/internal/storage"
	"fmt"
)

func ListTasks(status string) {
	valid := map[string]bool{
		"todo":        true,
		"done":        true,
		"in-progress": true,
	}

	if status != "" && !valid[status] {
		fmt.Println("Unknown status!")
		return
	}

	found := false
	tasks := storage.LoadTasks()

	for _, t := range tasks {
		if status == "" || t.Status == status {
			fmt.Printf("%d: %s [%s]\n", t.ID, t.Description, t.Status)
			found = true
		}
	}

	if !found {
		fmt.Println("No tasks found")
	}
}
