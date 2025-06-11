package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerPort   int
	DatabaseUrl  string
	JWTSecertKey string
}

func LoadConfig() *Config {

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	serverPort, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port value from env variable PORT: %s", portStr)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("Unable to import database_url from env variable")
	}

	jwtSecertKey := os.Getenv("JWT_SECERT_KEY")
	if jwtSecertKey == "" {
		log.Fatal("unable to jwt_secert_key from env variable")
	}

	return &Config{
		ServerPort:   serverPort,
		DatabaseUrl:  databaseURL,
		JWTSecertKey: jwtSecertKey,
	}

}
