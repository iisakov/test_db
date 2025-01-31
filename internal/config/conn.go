package config

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"test_bd/internal/api"
)

func newConn() *sqlx.DB {
	// Подключаемся к базе данных SQLite
	db, err := sqlx.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewRouter() *gin.Engine {
	// Инициализация Gin
	router := gin.Default()

	db := newConn()

	CreateTable(db)

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Маршруты
	router.GET("/tasks", api.GetTasks)          // Получить все задачи
	router.GET("/tasks/:id", api.GetTask)       // Получить задачу по ID
	router.POST("/tasks", api.CreateTask)       // Создать новую задачу
	router.PUT("/tasks/:id", api.UpdateTask)    // Обновить задачу
	router.DELETE("/tasks/:id", api.DeleteTask) // Удалить задачу

	// Добавляем маршрут для Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
