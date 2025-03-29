package repository

import (
	"context"
	"register-profile-service/internal/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUser(UserID int64) (interface{}, error)
	AddUser(userData models.User) (string, error)
}

type Repository struct {
	User
}

func NewRepository(ctx context.Context, db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(ctx, db),
	}
}
