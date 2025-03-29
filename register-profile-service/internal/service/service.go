package service

import (
	"context"
	"net/http"
	"register-profile-service/internal/repository"
)

type Service struct {
}

type UserService interface {
	SingUp(w http.ResponseWriter, r *http.Request)
	SingIn(w http.ResponseWriter, r *http.Request)
}

func NewService(ctx context.Context, repository *repository.Repository) *Service {

}
