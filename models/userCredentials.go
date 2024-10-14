package models

// UserCredentials holds the user input for login
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
