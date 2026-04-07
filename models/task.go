package models

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title" validate:"required"`
	Done  bool   `json:"done"`
}
