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
	connStr := "postgresql://godbtest_user:q0cT0UalkNc5zLyHoUs1YrIB6iGCDkWV@dpg-cvgnmnofnakc73fg6h6g-a.oregon-postgres.render.com:5432/godbtest_v3"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
