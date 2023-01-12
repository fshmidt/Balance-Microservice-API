package service

import (
	balance_api "balance-api"
	"balance-api/pkg/repository"
)

type TransactionsService struct {
	repo repository.Transactions
}

func NewTransactionsService(repo repository.Transactions) *TransactionsService {
	return &TransactionsService{repo: repo}
}

func (s *TransactionsService) GetTransactions(id, flag int) ([]balance_api.Transactions, error) {
	return s.repo.GetTransactions(id, flag)
}
