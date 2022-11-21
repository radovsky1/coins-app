package core

type Transfer struct {
	Id            int    `json:"-" db:"id"`
	FromAccountID int    `json:"from_account_id" db:"from_account_id"`
	ToAccountID   int    `json:"to_account_id" db:"to_account_id"`
	Amount        int64  `json:"amount" db:"amount"`
	Currency      string `json:"currency" db:"currency"`
}
