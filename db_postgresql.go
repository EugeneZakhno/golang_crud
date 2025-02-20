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
	connStr := "postgresql://godbtest_user:944sBzadpIXTdF9yGaxg5TFTQyTeaWxY@dpg-curoet52ng1s73dkune0-a.oregon-postgres.render.com/godbtest_21mb_of6a"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
