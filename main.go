package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// Структура для представления пользователя
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Глобальная переменная для хранения базы данных
var db *sql.DB

// Инициализация базы данных
func initDB() {
	var err error
	connStr := "postgresql://godbtest_user:lUDEQDsf2MrpRu80RajTBSOG70RNBcY4@dpg-cu74g1q3esus73fg1beg-a.oregon-postgres.render.com/godbtest_21mb"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

// Получение всех пользователей
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Создание нового пользователя
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.QueryRow("INSERT INTO users(name, age) VALUES($1, $2) RETURNING id", user.Name, user.Age).Scan(&user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Получение пользователя по ID
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	err := db.QueryRow("SELECT id, name, age FROM users WHERE id = $1", params["id"]).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Обновление пользователя
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", user.Name, user.Age, params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Удаление пользователя
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := db.Exec("DELETE FROM users WHERE id = $1", params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Сервер запущен на порту :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
