package service

import (
	"TaskTracker/internal/config"
	"TaskTracker/internal/model"
	"TaskTracker/internal/storage"
	"TaskTracker/internal/utils"
	"fmt"
	"time"
)

func AddTask(args []string) {
	if len(args) < 3 {
		fmt.Println("Not enough arguments")
		return
	}

	now := time.Now().Format(config.Format)

	tasks := storage.LoadTasks()
	newTask := model.Task{
		ID:          utils.GetNextID(tasks),
		Description: args[2],
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, newTask)

	err := storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:  ", err)
		return
	}

	fmt.Println("Task added!♥")
	fmt.Printf("%s (ID: %d)\n", args[2], newTask.ID)
}
