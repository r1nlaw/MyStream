package handler

import (
	"register-profile-service/internal/service"

	"github.com/rs/cors"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() {

	corsPolice := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

}
