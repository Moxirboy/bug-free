package handler

import "golang-project-template/internal/service/usecase"

type Handler struct {
	service *usecase.Crud
}

func NewHandler(service *usecase.Crud) *Handler {
	return &Handler{service: service}
}
