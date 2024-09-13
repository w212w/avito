package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(connectionString string) {
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Не удалось установить соединение с базой данных: %v", err)
	}
	log.Println("Соединение с базой данных успешно установлено")
}

func GetDB() *sql.DB {
	return db
}
