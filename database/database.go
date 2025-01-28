package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("hostDB"),
		os.Getenv("portDB"),
		os.Getenv("userDB"),
		os.Getenv("passwordDB"),
		os.Getenv("ejNameDB"),
		os.Getenv("sslModeDB"))

	// Открытие соединения
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	}
	defer DB.Close()

	// Проверка соединения
	if err = DB.Ping(); err != nil {
		fmt.Printf("Не удалось установить соединение: %v\n", err)
	} else {
		fmt.Println("DB connected")
	}
}
