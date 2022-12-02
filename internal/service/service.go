package service

import (
	"coins-app/internal/core"
	"coins-app/internal/es"
	"coins-app/internal/service/webapi"
	"coins-app/internal/storage"
)

//go:generate mockgen -source=service.go -destination=mocks/mock_service.go

type Authorization interface {
	CreateUser(user core.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Coin interface {
	GetCoinPrices(symbol string) ([]core.SymbolPrice, error)
}

type Service struct {
	Authorization
	Account
	Transfer
	Coin
}

func NewService(storages *storage.Storage, webapi *webapi.BinanceWebAPI, msgBroker *es.MessageBroker) *Service {
	return &Service{
		Authorization: NewAuthService(storages.Authorization),
		Account:       NewAccountService(storages.Account, msgBroker.Account),
		Transfer:      NewTransferService(storages.Transfer, storages.Account, msgBroker.Transfer),
		Coin:          NewCoinService(webapi),
	}
}
