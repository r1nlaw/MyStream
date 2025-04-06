package handler

import (
	"net/http"
	"register-profile-service/internal/service"

	"github.com/rs/cors"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() http.Handler {
	router := http.NewServeMux()
	
	corsPolicy := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            false,
	})

	router.HandleFunc("/MS/signIn", h.service.SignIn)
	router.HandleFunc("/MS/signUp", h.service.SignUp)

	handler := corsPolicy.Handler(router)

	return handler
}
