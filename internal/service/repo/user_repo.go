package repo

import (
	"database/sql"
	"bug-free/internal/model"
)

type postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *postgres {
	return &postgres{db: db}
}

type PostgresRepo interface {
	Create(user model.User) (int, error)
	ReadById(id model.UserGetBYId) (string,error)
	Update(user model.User) (int,error)
	Delete(user model.User) (int,error)
}
