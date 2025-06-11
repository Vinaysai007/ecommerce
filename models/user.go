package models

import (
	"time"
)

type User struct {
	ID           int       `json:"user_id"`
	Email        string    `json:"user_email"`
	PasswordHash string    `json:"-"`
	FirstName    string    `json:"user_firstname"`
	LastName     string    `json:"user_secondname"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
