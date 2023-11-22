package usecase

import (
	"bug-free/internal/model"
	"bug-free/internal/service/repo"
)

type User struct {
	repo repo.Postgres_
}

func NewCrud(repo repo.Postgres_) *User {
	return &User{repo: repo}
}

type UserUse interface {
	Create(user model.User) (int, error)
	Read() []string
	Update() error
	Delete() error
}
