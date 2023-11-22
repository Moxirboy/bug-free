package usecase

import (
	"bug-free/internal/model"
	"bug-free/internal/service/repo"
)

type User struct {
	repo repo.PostgresRepo
}

func NewCrud(repo repo.PostgresRepo) *User {
	return &User{repo: repo}
}

type UserUse interface {
	Create(user model.User) (int, error)
	Read() []string
	Update() error
	Delete() error
}
