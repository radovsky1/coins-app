package psql

import (
	"coins-app/internal/core"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) CreateAccount(account core.Account) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_id, balance, currency) values ($1, $2, $3) RETURNING id", accountsTable)

	row := r.db.QueryRow(query, account.UserId, account.Balance, account.Currency)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AccountPostgres) GetAccountById(accountId int) (core.Account, error) {
	var account core.Account

	query := fmt.Sprintf("SELECT id, user_id, balance, currency FROM %s WHERE id=$1", accountsTable)

	if err := r.db.QueryRow(query, accountId).Scan(&account.Id, &account.UserId, &account.Balance, &account.Currency); err != nil {
		return core.Account{}, err
	}

	return account, nil
}

func (r *AccountPostgres) GetAccounts(userId int) ([]core.Account, error) {
	var accounts []core.Account

	query := fmt.Sprintf("SELECT id, user_id, balance, currency FROM %s WHERE user_id=$1", accountsTable)

	if err := r.db.Select(&accounts, query, userId); err != nil {
		return []core.Account{}, err
	}

	return accounts, nil
}

func (r *AccountPostgres) UpdateAccount(account core.Account) error {
	query := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id=$2", accountsTable)

	_, err := r.db.Exec(query, account.Balance, account.Id)

	return err
}
