package utils

import (
	"log"
	"time"

	"github.com/Vinaysai007/ecom/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(useremail string) (string, error) {
	cfg := config.LoadConfig()
	JWTSecertKey := []byte(cfg.JWTSecertKey)

	claims := &Claims{
		Email: useremail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecertKey)
	if err != nil {
		log.Printf("Error: failed to sign JWT token: %v", err)
		return "", err
	}
	return tokenString, nil

}

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
