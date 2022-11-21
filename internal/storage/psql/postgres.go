package psql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	// AccountsTable is the name of the account table
	accountsTable = "accounts"
	// UsersTable is the name of the user table
	usersTable = "users"
	// TransfersTable is the name of the transfer table
	transfersTable = "transfers"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
