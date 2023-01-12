package repository

import (
	balance_api "balance-api"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransactionsPostgres struct {
	db *sqlx.DB
}

func NewTransactionsPostgres(db *sqlx.DB) *TransactionsPostgres {
	return &TransactionsPostgres{db: db}
}

func (r *TransactionsPostgres) GetTransactions(id, flag int) ([]balance_api.Transactions, error) {
	var lists []balance_api.Transactions
	var tail string
	if flag == 1 {
		tail = "ORDER BY netto DESC"
	}
	query := fmt.Sprintf(`SELECT id, user_id, netto, cashflow, source_or_purpose, transTime FROM %s WHERE user_id = $1 %s`, transactionsTable, tail)
	err := r.db.Select(&lists, query, id)

	return lists, err
}
