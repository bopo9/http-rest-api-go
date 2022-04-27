package repository

import (
	"github.com/jmoiron/sqlx"
	socialApp "restApiGo"
)

type Authorization interface {
	CreateUser(user socialApp.User) (int, error)
	GetUser(username, password string) (socialApp.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
