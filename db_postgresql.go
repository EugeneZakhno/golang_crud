package main

import (
	"database/sql"
	"log"
)

// Глобальная переменная для хранения базы данных
var db *sql.DB

// Инициализация базы данных
func InitDB() {
	var err error
	connStr := "postgresql://godbtest_8g85_user:gHNGuGjrVDcjlM9TdIDm8nTFFguj1QnU@dpg-d0q4hrvdiees738o58d0-a.oregon-postgres.render.com/godbtest_8g85"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
