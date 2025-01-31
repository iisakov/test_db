package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
	myDb "test_bd/internal/db"
	"test_bd/internal/model"
)

// GetTasks возвращает список всех задач
// @Summary Получить все задачи
// @Description Возвращает список всех задач
// @Tags tasks
// @Produce json
// @Success 200 {array} model.Task
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	rows, err := db.Query(myDb.GetTasks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() { _ = rows.Close() }()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTask возвращает задачу по ID
// @Summary Получить задачу по ID
// @Description Возвращает задачу по её ID
// @Tags tasks
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} model.Task
// @Failure 404 {string} string "Задача не найдена"
// @Router /tasks/{id} [get]
func GetTask(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	id := c.Param("id")
	var task model.Task
	err := db.QueryRow(myDb.GetTask, id).Scan(&task.ID, &task.Title, &task.Done)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CreateTask создает новую задачу
// @Summary Создать задачу
// @Description Создает новую задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Данные задачи"
// @Success 201 {object} model.Task
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	result, err := db.Exec(myDb.CreateTask, newTask.Title, newTask.Done)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	newTask.ID = int(id)

	c.JSON(http.StatusCreated, newTask)
}

// UpdateTask обновляет существующую задачу
// @Summary Обновить задачу
// @Description Обновляет существующую задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param task body model.Task true "Обновленные данные задачи"
// @Success 200 {object} model.Task
// @Failure 404 {string} string "Задача не найдена"
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	id := c.Param("id")
	var updatedTask model.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	_, err := db.Exec(myDb.UpdateTask, updatedTask.Title, updatedTask.Done, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedTask.ID, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask удаляет задачу по ID
// @Summary Удалить задачу
// @Description Удаляет задачу по её ID
// @Tags tasks
// @Param id path int true "ID задачи"
// @Success 204 "No Content"
// @Failure 404 {string} string "Задача не найдена"
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	id := c.Param("id")
	_, err := db.Exec(myDb.DeleteTask, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
