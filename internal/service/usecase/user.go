package usecase

import (
	"bug-free/internal/model"
	"bug-free/internal/service/repo"
)

type User struct {
	repo repo.PostgresRepo
}

func NewUser(repo repo.PostgresRepo) *User {
	return &User{repo: repo}
}

type UserUse interface {
	Create(user model.User) (int, error)
	Read(id model.UserGetBYId) (string,error)
	Update(user model.User) (int ,error)
	Delete(user model.User) (int,error)
}
