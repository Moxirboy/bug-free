package handler

import "bug-free/internal/service/usecase"

type Handler struct {
	service *usecase.User
}

func NewHandler(service *usecase.User) *Handler {
	return &Handler{service: service}
}
