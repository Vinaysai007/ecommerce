package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Vinaysai007/ecom/models"
	"github.com/Vinaysai007/ecom/utils"
)

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var UserInput struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"second_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&UserInput); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if UserInput.Email == "" || UserInput.Password == "" {
		http.Error(w, "Email and Password are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(UserInput.Password)
	if err != nil {
		http.Error(w, "Failed to hash the password", http.StatusInternalServerError)
		log.Printf("ERROR: Failed to hash password for %s: %v", UserInput.Email, err)
		return
	}

	user := models.User{
		Email:        UserInput.Email,
		PasswordHash: hashedPassword,
		FirstName:    UserInput.FirstName,
		LastName:     UserInput.LastName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query :=
		`INSERT INTO users(id,email,password_hash,first_name,last_name,created_at,updated_at)
		VALUES($1,$2,$3,$4,$5,$6) 
		RETURNING id;`

	var userID int
	err = uh.DB.QueryRow(query, user.Email, user.PasswordHash, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt).Scan(&userID)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
			http.Error(w, "User with this email already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		log.Printf("Error: Failed to insert user %s into database: %v", UserInput.Email, err)
		return
	}

	user.ID = userID
	user.PasswordHash = ""
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	log.Printf("user registered successfully")
}
