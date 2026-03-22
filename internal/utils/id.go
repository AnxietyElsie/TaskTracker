package utils

import (
	"TaskTracker/internal/model"
)

func GetNextID(tasks []model.Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}
