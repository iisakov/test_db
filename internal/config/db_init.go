package config

import (
	"github.com/jmoiron/sqlx"
	"log"
)

// CreateTable создает таблицу задач, если она не существует
func CreateTable(db *sqlx.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY,
        title TEXT NOT NULL,
        done BOOLEAN NOT NULL DEFAULT FALSE
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
