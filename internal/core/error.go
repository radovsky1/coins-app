package core

import (
	"errors"
)

var (
	ErrInsufficientFunds    = errors.New("insufficient funds")
	ErrAccountNotFound      = errors.New("account not found")
	ErrUserNotFound         = errors.New("user not found")
	ErrTransferNotFound     = errors.New("transfer not found")
	ErrAmountMustBePositive = errors.New("amount must be positive")
	ErrSameAccount          = errors.New("same account")
)
