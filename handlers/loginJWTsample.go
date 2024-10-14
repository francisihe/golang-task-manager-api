package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/francisihe/golang-task-manager-api/models"
)

// Secret key (same as in middleware)
var SecretKey = []byte("my_secret_key")

// Login endpoint to generate JWT
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.UserCredentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate user credentials (you would normally check this against the DB)
	if credentials.Username != "admin" || credentials.Password != "password" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create a new JWT token
	claims := &jwt.StandardClaims{
		Subject:   "user_id", // Normally use the user ID from DB
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
