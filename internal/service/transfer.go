package service

import (
	"coins-app/internal/core"
	"coins-app/internal/storage"
)

type TransferService struct {
	TransferRepo storage.Transfer
	AccountRepo  storage.Account
}

func NewTransferService(transferRepo storage.Transfer, accountRepo storage.Account) *TransferService {
	return &TransferService{TransferRepo: transferRepo, AccountRepo: accountRepo}
}

func (s *TransferService) CreateTransfer(transfer core.Transfer) (int, error) {
	if err := s.validTransfer(transfer); err != nil {
		return 0, err
	}
	transferId, err := s.TransferRepo.CreateTransfer(transfer)
	if err != nil {
		return 0, err
	}
	return transferId, nil
}

func (s *TransferService) validAccount(accountId int) (core.Account, error) {
	account, err := s.AccountRepo.GetAccountById(accountId)
	if err != nil {
		return core.Account{}, err
	}
	return account, nil
}

func (s *TransferService) validTransfer(transfer core.Transfer) error {
	if transfer.Amount <= 0 {
		return core.ErrAmountMustBePositive
	}
	if transfer.FromAccountID == transfer.ToAccountID {
		return core.ErrSameAccount
	}

	_, err := s.validAccount(transfer.FromAccountID)
	if err != nil {
		return err
	}

	_, err = s.validAccount(transfer.ToAccountID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TransferService) GetTransferById(transferId int) (core.Transfer, error) {
	transfer, err := s.TransferRepo.GetTransferById(transferId)
	if err != nil {
		return core.Transfer{}, err
	}
	return transfer, nil
}

func (s *TransferService) GetTransfers(userId int) ([]core.Transfer, error) {
	transfers, err := s.TransferRepo.GetTransfers(userId)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}
