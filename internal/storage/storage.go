package storage

import (
	"TaskTracker/internal/config"
	"TaskTracker/internal/model"
	"encoding/json"
	"fmt"
	"os"
)

func LoadTasks() []model.Task {
	data, err := os.ReadFile(config.TasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error task file is not exist: ", err)
			return []model.Task{}
		}
		fmt.Println("Error reading tasks file: ", err)
		return []model.Task{}
	}

	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error JSON parsing: ", err)
		return []model.Task{}
	}

	return tasks
}

func SaveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error formating JSON: ", err)
		return err
	}
	err = os.WriteFile(config.TasksFile, data, 0644)
	if err != nil {
		fmt.Println("Error: failed to save task", err)
		return err
	}
	return nil
}
