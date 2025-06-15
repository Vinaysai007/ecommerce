package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Vinaysai007/ecom/config"
	db "github.com/Vinaysai007/ecom/database"
	"github.com/Vinaysai007/ecom/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config.LoadConfig()
	db.InitDB(cfg.DatabaseUrl)
	defer db.DB.Close()

	db.CreateTableUsers()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the E-commerece Backend API!")
	})

	userHandler := handlers.NewUserHandler(db.DB)
	mux.HandleFunc("/user/register", userHandler.RegisterUser)

	port := ":" + strconv.Itoa(cfg.ServerPort)
	fmt.Printf("Server starting on port: %s...\n", port)
	log.Fatal(http.ListenAndServe(port, mux))

}
