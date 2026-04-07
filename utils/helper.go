package utils

import "stability-test-task-api/models"

func GetMaxID(tasks []models.Task) int {
	var maxID = 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID
}
