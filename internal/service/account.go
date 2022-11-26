package service

import (
	"coins-app/internal/core"
	"coins-app/internal/storage"
	"coins-app/util"
)

type AccountService struct {
	AccountRepo storage.Account
}

func NewAccountService(accountRepo storage.Account) *AccountService {
	return &AccountService{AccountRepo: accountRepo}
}

func (s *AccountService) CreateAccount(account core.Account) (int, error) {
	if !util.IsSupportedCoin(account.Currency) {
		return 0, core.ErrUnsupportedCurrency
	}
	accountId, err := s.AccountRepo.CreateAccount(account)
	if err != nil {
		return 0, err
	}
	return accountId, nil
}

func (s *AccountService) GetAccountById(accountId int) (core.Account, error) {
	account, err := s.AccountRepo.GetAccountById(accountId)
	if err != nil {
		return core.Account{}, err
	}
	return account, nil
}

func (s *AccountService) GetAccounts(userId int) ([]core.Account, error) {
	accounts, err := s.AccountRepo.GetAccounts(userId)
	if err != nil {
		return []core.Account{}, err
	}
	return accounts, nil
}

func (s *AccountService) UpdateAccount(account core.Account) error {
	err := s.AccountRepo.UpdateAccount(account)
	if err != nil {
		return err
	}
	return nil
}
