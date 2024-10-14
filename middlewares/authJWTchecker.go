package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Secret key for signing the JWTs (use a secure way to store it in production)
var SecretKey = []byte("my_secret_key")

// Middleware to validate JWT in request header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Split the header value to get the token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check if the token's signing method is HMAC and return the secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return SecretKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}
