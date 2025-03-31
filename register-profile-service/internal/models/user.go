package models

import "time"

type User struct {
	ID           int64     `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	IsVerified   bool      `json:"is_verified" db:"is_verified"`
	Created_at   time.Time `json:"created_at" db:"created_at"`
	Updated_at   time.Time `json:"updated_at" db:"updated_at"`
}
