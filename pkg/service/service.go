package service

import (
	balance_api "balance-api"
	"balance-api/pkg/repository"
)

type Balance interface {
	GetBalance(id int) (int, error)
	UpdateBalance(id int, input balance_api.UpdateBalance, sourceOrPurpose string) error
}

type Purchase interface {
	UpdatePurchase(id int, srv string) error
}

type Transactions interface {
	GetTransactions(id, flag int) ([]balance_api.Transactions, error)
}

type Service struct {
	Balance
	Purchase
	Transactions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Balance:      NewBalanceService(repos.Balance),
		Purchase:     NewPurchaseService(repos.Purchase),
		Transactions: NewTransactionsService(repos.Transactions),
	}
}
