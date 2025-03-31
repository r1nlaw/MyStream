package repository

import (
	"context"
	"database/sql"
	"errors"
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

func (u *UserDB) GetUser(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	var result models.User
	err := u.postgres.GetContext(ctx, &result, query, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &result, nil
}

func (u *UserDB) AddUser(ctx context.Context, userData models.User) error {
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`
	_, err := u.postgres.ExecContext(ctx, query, userData.Username, userData.Email, userData.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to add user %w", err)
	}
	return nil
}
