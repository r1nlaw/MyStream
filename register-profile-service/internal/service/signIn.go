package service

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type SignUpRepository struct {
	Db *sqlx.DB
}

func NewSignUpRepository(db *sqlx.DB) *SignUpRepository {
	return &SignUpRepository{Db: db}
}

func (c *SignUpRepository) SignUp(w http.ResponseWriter, r *http.Request) {

}
