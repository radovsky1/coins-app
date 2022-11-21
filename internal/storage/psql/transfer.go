package psql

import (
	"coins-app/internal/core"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransferPostgres struct {
	db *sqlx.DB
}

func NewTransferPostgres(db *sqlx.DB) *TransferPostgres {
	return &TransferPostgres{db: db}
}

func (t *TransferPostgres) CreateTransfer(transfer core.Transfer) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var transferId int
	createTransferQuery := fmt.Sprintf("INSERT INTO %s (from_account_id, to_account_id, amount, currency) values ($1, $2, $3, $4) RETURNING id", transfersTable)

	row := tx.QueryRow(createTransferQuery, transfer.FromAccountID, transfer.ToAccountID, transfer.Amount, transfer.Currency)
	if err := row.Scan(&transferId); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	updateFromAccountQuery := fmt.Sprintf("UPDATE %s SET balance=balance-$1 WHERE id=$2", accountsTable)
	_, err = tx.Exec(updateFromAccountQuery, transfer.Amount, transfer.FromAccountID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	updateToAccountQuery := fmt.Sprintf("UPDATE %s SET balance=balance+$1 WHERE id=$2", accountsTable)
	_, err = tx.Exec(updateToAccountQuery, transfer.Amount, transfer.ToAccountID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return transferId, nil
}

func (t *TransferPostgres) GetTransferById(transferId int) (core.Transfer, error) {
	var transfer core.Transfer

	query := fmt.Sprintf("SELECT id, from_account_id, to_account_id, amount, currency FROM %s WHERE id=$1", transfersTable)

	if err := t.db.QueryRow(query, transferId).Scan(&transfer.Id, &transfer.FromAccountID, &transfer.ToAccountID, &transfer.Amount, &transfer.Currency); err != nil {
		return core.Transfer{}, err
	}

	return transfer, nil
}

func (t *TransferPostgres) GetTransfers(userId int) ([]core.Transfer, error) {
	var transfers []core.Transfer

	query := fmt.Sprintf("SELECT id, from_account_id, to_account_id, amount, currency FROM %s WHERE from_account_id=$1", transfersTable)

	if err := t.db.Select(&transfers, query, userId); err != nil {
		return []core.Transfer{}, err
	}

	return transfers, nil
}
