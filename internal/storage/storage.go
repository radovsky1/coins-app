package storage

import (
	"coins-app/internal/core"
	"coins-app/internal/storage/psql"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user core.User) (int, error)
	GetUser(username, password string) (core.User, error)
}

type Account interface {
	CreateAccount(account core.Account) (int, error)
	GetAccountById(accountId int) (core.Account, error)
	GetAccounts(userId int) ([]core.Account, error)
	UpdateAccount(account core.Account) error
}

type Transfer interface {
	CreateTransfer(transfer core.Transfer) (int, error)
	GetTransferById(transferId int) (core.Transfer, error)
	GetTransfers(userId int) ([]core.Transfer, error)
}

type Storage struct {
	Authorization
	Account
	Transfer
}

func NewStoragePostgres(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: psql.NewAuthPostgres(db),
		Account:       psql.NewAccountPostgres(db),
		Transfer:      psql.NewTransferPostgres(db),
	}
}
