package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(databaseURL string) {
	var err error
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect the database, provide correct details")
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database unreachable:", err)
	}
	log.Println("database connected successfully")
}

func CreateTableUsers() {
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("unable to insert the user details into the table users in database", err)
	}
}
