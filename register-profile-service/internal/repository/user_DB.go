package repository

import (
	"context"
	"fmt"
	"register-profile-service/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	ctx      context.Context
	postgres *sqlx.DB
}

func NewUserPostgres(ctx context.Context, db *sqlx.DB) *UserDB {
	return &UserDB{ctx: ctx, postgres: db}
}

func (u *UserDB) GetUser(ctx context.Context, email string) (interface{}, error) {
	query := `SELECT * FROM users WHERE email=$1`
	var result interface{}
	err := u.postgres.Select(&result, query)
	if err != nil {
		return "", fmt.Errorf("failed to get user %w", err)
	}
	return result, nil
}

func (u *UserDB) AddUser(ctx context.Context, userData models.User) error {
	quary := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`
	var newID int
	err := u.postgres.QueryRow(quary, userData.Username, userData.Email, userData.PasswordHash).Scan(&newID)
	if err != nil {
		return fmt.Errorf("failed to add user %w", err)
	}
	return fmt.Errorf("user add successfully")
}
