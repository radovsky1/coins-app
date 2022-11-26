package psql

import (
	"coins-app/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestTransferPsql_CreateTransfer(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	r := NewTransferPostgres(db)
	tests := []struct {
		name    string
		mock    func()
		input   core.Transfer
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO transfers").
					WithArgs(1, 2, 100, "USD").WillReturnRows(rows)

				mock.ExpectExec("UPDATE accounts SET").
					WithArgs(100, 1).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("UPDATE accounts SET").
					WithArgs(100, 2).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			input: core.Transfer{
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        100,
				Currency:      "USD",
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			mock: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO transfers").
					WithArgs(1, 2, 100, "").WillReturnRows(rows)

				mock.ExpectExec("UPDATE accounts SET").
					WithArgs(100, 1).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("UPDATE accounts SET").
					WithArgs(100, 2).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			input: core.Transfer{
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        100,
				Currency:      "",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got, err := r.CreateTransfer(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestTransferPsql_GetTransferById(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	r := NewTransferPostgres(db)
	tests := []struct {
		name    string
		mock    func()
		input   int
		want    core.Transfer
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "from_account_id", "to_account_id", "amount", "currency"}).
					AddRow(1, 1, 2, 100, "USD")
				mock.ExpectQuery("SELECT (.+) FROM transfers").
					WithArgs(1).WillReturnRows(rows)
			},
			input: 1,
			want: core.Transfer{
				Id:            1,
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        100,
				Currency:      "USD",
			},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "from_account_id", "to_account_id", "amount", "currency"})
				mock.ExpectQuery("SELECT (.+) FROM transfers").
					WithArgs(1).WillReturnRows(rows)
			},
			input:   1,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got, err := r.GetTransferById(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestTransferPsql_GetTransfers(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	r := NewTransferPostgres(db)
	tests := []struct {
		name    string
		mock    func()
		input   int
		want    []core.Transfer
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "from_account_id", "to_account_id", "amount", "currency"}).
					AddRow(1, 1, 2, 100, "USD")
				mock.ExpectQuery("SELECT (.+) FROM transfers").
					WithArgs(1).WillReturnRows(rows)
			},
			input: 1,
			want: []core.Transfer{
				{
					Id:            1,
					FromAccountID: 1,
					ToAccountID:   2,
					Amount:        100,
					Currency:      "USD",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got, err := r.GetTransfers(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
