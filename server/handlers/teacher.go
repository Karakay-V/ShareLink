package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Karakay-V/ShareLink/server/config"
	"github.com/Karakay-V/ShareLink/server/db"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func TeacherLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var id int
	var hash string
	err := db.DB.QueryRow("SELECT id, password_hash FROM teachers WHERE email = ?", req.Email).Scan(&id, &hash)
	if err == sql.ErrNoRows {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	// compare password
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)) != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// success → generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   req.Email,
		"role":    "teacher",
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString(config.JwtSecret)

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func RegisterTeacher(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	_, err = db.DB.Exec(
		`INSERT INTO teachers (first_name, middle_name, last_name, email, password_hash) 
		 VALUES (?, ?, ?, ?, ?)`,
		req.FirstName, req.MiddleName, req.LastName, req.Email, string(hash),
	)

	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "teacher created"})
}
