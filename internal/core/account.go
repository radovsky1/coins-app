package core

type Account struct {
	Id       int    `json:"-" db:"id"`
	UserId   int    `json:"user_id" db:"user_id"`
	Balance  int64  `json:"balance" db:"balance"`
	Currency string `json:"currency" db:"currency"`
}

func (a *Account) CanTransfer(amount int64) bool {
	return a.Balance >= amount
}

func (a *Account) Transfer(amount int64) {
	a.Balance -= amount
}

func (a *Account) Add(amount int64) {
	a.Balance += amount
}
