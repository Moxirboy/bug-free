package usecase

import (
	"bug-free/internal/model"
	"bug-free/internal/service/repo"
)

type Crud struct {
	repo repo.Postgres_
}

func NewCrud(repo repo.Postgres_) *Crud {
	return &Crud{repo: repo}
}

type crud_ interface {
	Create(user model.User) (int, error)
	Read() []string
	Update() error
	Delete() error
}