package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var PsqlDB *sql.DB

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

	PsqlDB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	}

	if err = PsqlDB.Ping(); err != nil {
		fmt.Printf("Pinging is failed. Error: %v\n", err)
	} else {
		fmt.Println("DB connected")
	}
}
