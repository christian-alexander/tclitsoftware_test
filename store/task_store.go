package store

import "stability-test-task-api/models"

var Tasks = []models.Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build API", Done: false},
}

func GetAllTasks() []models.Task {
	return Tasks
}

func GetTaskByID(id int) *models.Task {
	// returnan untuk tasks, jangan pakai copy an dari tasks, return pointer tasknya
	// return copyan dari tasks hanya cocok untuk data return yang read only
	// kalau sewaktu waktu porgram ini berkembang lebih kompleks dan get task dipakai di module lain
	// dan ada yang mutate hasil GetTaskById, maka instance Tasks nya tidak berubah karena yang direturn copyan
	for i := range Tasks {
		if Tasks[i].ID == id {
			return Tasks
		}
	}
	return nil
}

func AddTask(task models.Task) {
	Tasks = append(Tasks, task)
}

func DeleteTask(id int) {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
		}
	}
}
