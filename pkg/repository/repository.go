package repository

import (
	balance_api "balance-api"
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	GetBalance(id int) (int, error)
	UpdateBalance(id int, input balance_api.UpdateBalance, sourceOrPurpose string) error
	CreateBalance(id int, input balance_api.UpdateBalance) error
}

type Purchase interface {
	UpdatePurchase(id int, srv string) error
}

type Transactions interface {
	GetTransactions(id, flag int) ([]balance_api.Transactions, error)
}

type Repository struct {
	Balance
	Purchase
	Transactions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Balance:      NewBalancePostgres(db),
		Purchase:     NewPurchasePostgres(db),
		Transactions: NewTransactionsPostgres(db),
	}
}
