package psql

import (
	"coins-app/internal/core"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user core.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password) values ($1, $2, $3) RETURNING id", usersTable)

	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (core.User, error) {
	var user core.User

	query := fmt.Sprintf("SELECT id, username, password FROM %s WHERE username=$1 AND password=$2", usersTable)

	if err := a.db.Get(&user, query, username, password); err != nil {
		return core.User{}, err
	}

	return user, nil
}
