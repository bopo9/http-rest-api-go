package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	socialApp "restApiGo"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user socialApp.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, surname, password_hash) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.SecondName, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (socialApp.User, error) {
	var user socialApp.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
