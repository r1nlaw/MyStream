package service

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type SignInRepository struct {
	Db *sqlx.DB
}

func NewSignInRepository(db *sqlx.DB) *SignInRepository {
	return &SignInRepository{Db: db}
}

func (c *SignInRepository) SignIn(w http.ResponseWriter, r *http.Request) {

}
