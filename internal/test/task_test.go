package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"test_bd/internal/db"
	"test_bd/internal/test/expected_result"

	"net/http"
	"net/http/httptest"
	"test_bd/internal/api"
	"testing"
)

// TestGetTasksSuccess тестирует успешный запрос
func TestGetTasksSuccess(t *testing.T) {
	// Создаем mock-объект для базы данных с помощью db-dog
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() { _ = mockDB.Close() }()

	// Настраиваем ожидаемый запрос и возвращаемые данные
	rows := sqlmock.NewRows([]string{"id", "title", "done"}).
		AddRow(1, "Task 1", false).
		AddRow(2, "Task 2", true)

	mock.ExpectQuery(db.GetTasks).WillReturnRows(rows)

	// Создаем тестовый HTTP-контекст Gin
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("db", mockDB)

	// Вызываем тестируемый метод
	api.GetTasks(c)

	// Проверяем результат
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expected_result.GetTask, w.Body.String())

	// Проверяем, что все ожидания mock-объекта были выполнены
	assert.NoError(t, mock.ExpectationsWereMet())
}
