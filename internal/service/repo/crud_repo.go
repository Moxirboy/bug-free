package repo

import (
	"database/sql"
	"golang-project-template/internal/model"
)

type postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *postgres {
	return &postgres{db: db}
}

type Postgres_ interface {
	Create(user model.User) (int, error)
	Read() []string
	Update() error
	Delete() error
}
