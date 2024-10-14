package middleware

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var (
	// Define rate limiter with 5 requests per second and a burst size of 10
	limiter = rate.NewLimiter(rate.Every(1*time.Second), 10)
	mu      sync.Mutex
)

// RateLimiting middleware to limit API calls
func RateLimiting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		// Check if the request is allowed
		if !limiter.Allow() {
			http.Error(w, "Too many requests, please try again later.", http.StatusTooManyRequests)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}
