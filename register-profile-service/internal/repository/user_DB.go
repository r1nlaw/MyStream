package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"register-profile-service/internal/models"
	"time"

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
	tx, err := u.postgres.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`
	_, err = u.postgres.ExecContext(ctx, query, userData.Username, userData.Email, userData.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to add user %w", err)
	}

	var userID int
	err = tx.QueryRowContext(ctx, `SELECT id FROM users WHERE email = $1`, userData.Email).Scan(&userID)
	if err != nil {
		return fmt.Errorf("failed to get user id: %w", err)
	}
	profile := models.Profile{
		UserID:    userID,
		Username:  userData.Username,
		AvatarURL: "",
		UserBIO:   "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	profileQuery := `INSERT INTO profiles (user_id, username, avatar_url, user_bio, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = tx.ExecContext(ctx, profileQuery, profile.UserID, profile.Username, profile.AvatarURL, profile.UserBIO, profile.CreatedAt, profile.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to add profile: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (u *UserDB) AddProfile(ctx context.Context, profile models.Profile) error {
	query := `INSERT INTO profiles (user_id, username, avatar_url, user_bio, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := u.postgres.ExecContext(ctx, query, profile.UserID, profile.Username, profile.AvatarURL, profile.UserBIO, profile.CreatedAt, profile.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to add profile: %w", err)
	}
	return nil
}

func (u *UserDB) AddSession(ctx context.Context, jwt models.JWTRequest) error {
	query := `INSERT INTO sessions (user_id, token, created_at, expires_at) VALUES ($1, $2, $3, $4)`

	_, err := u.postgres.ExecContext(ctx, query, jwt.UserID, jwt.Token, jwt.CreatedAt, jwt.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to add session: %w", err)
	}
	return nil
}
