package model

// Task представляет структуру задачи
// @Description Задача (To-Do) с уникальным идентификатором, названием и статусом выполнения.
type Task struct {
	ID    int    `json:"id" example:"1"`                // Уникальный идентификатор задачи
	Title string `json:"title" example:"Купить молоко"` // Название задачи
	Done  bool   `json:"done" example:"false"`          // Статус выполнения задачи
}
