package psql

import (
	"coins-app/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestAccountPsql_CreateAccount(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	r := NewAccountPostgres(db)
	tests := []struct {
		name    string
		mock    func()
		input   core.Account
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO accounts").
					WithArgs(1, 0, "USD").WillReturnRows(rows)
			},
			input: core.Account{
				UserId:   1,
				Balance:  0,
				Currency: "USD",
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO accounts").
					WithArgs(1, 0, "").WillReturnRows(rows)
			},
			input: core.Account{
				UserId:   1,
				Balance:  0,
				Currency: "",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got, err := r.CreateAccount(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestAccountPsql_GetAccount(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	r := NewAccountPostgres(db)
	tests := []struct {
		name    string
		mock    func()
		input   int
		want    core.Account
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "balance", "currency"}).AddRow(1, 1, 0, "USD")
				mock.ExpectQuery("SELECT id, user_id, balance, currency FROM accounts").
					WithArgs(1).WillReturnRows(rows)
			},
			input: 1,
			want: core.Account{
				Id:       1,
				UserId:   1,
				Balance:  0,
				Currency: "USD",
			},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "balance", "currency"})
				mock.ExpectQuery("SELECT id, user_id, balance, currency FROM accounts").
					WithArgs(1).WillReturnRows(rows)
			},
			input:   1,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got, err := r.GetAccountById(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestAccountPsql_GetAccountsByUserId(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	r := NewAccountPostgres(db)
	tests := []struct {
		name    string
		mock    func()
		input   int
		want    []core.Account
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "balance", "currency"}).
					AddRow(1, 1, 0, "USD")
				mock.ExpectQuery("SELECT id, user_id, balance, currency FROM accounts").
					WithArgs(1).WillReturnRows(rows)
			},
			input: 1,
			want: []core.Account{
				{
					Id:       1,
					UserId:   1,
					Balance:  0,
					Currency: "USD",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got, err := r.GetAccounts(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestAccountPsql_UpdateAccount(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	r := NewAccountPostgres(db)
	tests := []struct {
		name    string
		mock    func()
		input   core.Account
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectExec("UPDATE accounts SET").
					WithArgs(0, 1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			input: core.Account{
				Id:       1,
				UserId:   1,
				Balance:  0,
				Currency: "USD",
			},
		},
		{
			name: "Empty Fields",
			mock: func() {
				mock.ExpectExec("UPDATE accounts SET").
					WithArgs(1, 0, "", 1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			input: core.Account{
				Id:       1,
				UserId:   1,
				Balance:  0,
				Currency: "",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := r.UpdateAccount(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
