package store

import (
	"stability-test-task-api/models"
	"stability-test-task-api/utils"
)

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
			return &Tasks[i]
		}
	}
	return nil
}

func AddTask(task models.Task) {
	// improvement 1 : if acc, id pakai auto increment, karena umumnya id tidak diinput sendiri. Namun: ini perlu crosscheck terkait penggunaan apinya untuk apa, terkadang ada api yang khusus menerima inputan id juga karena kebutuhan sinkronisasi antar project. Tujuan auto increment agar ada jaminan id unik.

	var newId = utils.GetMaxID(Tasks) + 1
	task.ID = newId

	// append new task
	Tasks = append(Tasks, task)
}

func DeleteTask(id int) {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
		}
	}
}
