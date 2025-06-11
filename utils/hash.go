package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash the password: %v", err)
		return "", err
	}
	return string(hashedPass), err
}

func ComparePassHash(HashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))
	if err != nil {
		log.Printf("Password incorrect: %v", err)
	}
	return nil
}
