package models

type SignInRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
