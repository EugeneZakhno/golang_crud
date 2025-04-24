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
	connStr := "postgresql://godbtest_user:eC2t9gs2B5dqydPl8oev3Luju4GXucgB@dpg-d059m1p5pdvs73etokbg-a.oregon-postgres.render.com/godbtest_1wll"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
