package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	_ "test_bd/docs"
	"test_bd/internal/config"
)

// @title Task API
// @version 1.0
// @description Это API для управления списком задач (To-Do).
// @host localhost:8080
// @BasePath /
func main() {
	// Инициализация Gin
	gin.SetMode("release")
	router := config.NewRouter()
	defer func() { _ = router.Run(":8080") }()
}
