package service

import (
	"context"
	"encoding/json"
	"net/http"
	"register-profile-service/internal/models"
	"register-profile-service/internal/repository"
	"register-profile-service/internal/service/hash"
	"register-profile-service/internal/service/token"
)

type Service struct {
	repository *repository.Repository
	ctx        context.Context
	tokenMaker token.Maker
	hashUtil   hash.Hasher
}

type UserService interface {
	SingUp(w http.ResponseWriter, r *http.Request)
	SingIn(w http.ResponseWriter, r *http.Request)
}

func NewService(ctx context.Context, repository *repository.Repository, tokenMaker token.Maker, hashUtil hash.Hasher) *Service {
	return &Service{
		repository: repository,
		ctx:        ctx,
		tokenMaker: tokenMaker,
		hashUtil:   hashUtil,
	}
}

func (s *Service) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var request models.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "failed to parse request", http.StatusBadRequest)
		return
	}
	hashedPassword, err := s.hashUtil.HashPassword(request.PasswordHash)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: hashedPassword,
	}

	if err := s.repository.AddUser(s.ctx, user); err != nil {
		http.Error(w, "failed to add user: %v", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created successfully"})
}

func (s *Service) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var request models.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "failed to parse request", http.StatusBadRequest)
		return
	}

	userData, err := s.repository.GetUser(s.ctx, request.Email)
	if err != nil {
		http.Error(w, "invalid email or passwordq", http.StatusUnauthorized)
		return
	}

	if !s.hashUtil.CheckPassword(userData.PasswordHash, request.PasswordHash) {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := s.tokenMaker.CreateToken(userData.ID)
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "login successful", "token": token})
}
