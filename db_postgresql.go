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
	connStr := "postgresql://godbtest_user:lUDEQDsf2MrpRu80RajTBSOG70RNBcY4@dpg-cu74g1q3esus73fg1beg-a.oregon-postgres.render.com/godbtest_21mb"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
