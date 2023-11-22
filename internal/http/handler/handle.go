package handler

import "bug-free/internal/service/usecase"

type Handler struct {
	service *usecase.Crud
}

func NewHandler(service *usecase.Crud) *Handler {
	return &Handler{service: service}
}
