package main

import (
	"log"
	"net/http"
	"os"
	"tender-service/handlers"
	"tender-service/repository"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	dbConnStr := os.Getenv("POSTGRES_CONN")
	if dbConnStr == "" {
		log.Fatal("POSTGRES_CONN не задан")
	}
	log.Printf("Используем строку подключения: %s", dbConnStr)

	repository.InitDB(dbConnStr)

	if repository.GetDB() == nil {
		log.Fatal("Не удалось установить подключение к базе данных")
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/ping", handlers.PingHandler).Methods("GET")
	r.HandleFunc("/api/tenders", handlers.TenderListHandler).Methods("GET")
	r.HandleFunc("/api/tenders/my", handlers.GetUserTendersHandler).Methods("GET")
	r.HandleFunc("/api/tenders/new", handlers.NewTenderHandler).Methods("POST")
	r.HandleFunc("/api/tenders/status", handlers.UpdateTenderStatusHandler).Methods("PATCH")
	r.HandleFunc("/api/bids/new", handlers.NewBidHandler).Methods("POST")

	log.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", r)
}
