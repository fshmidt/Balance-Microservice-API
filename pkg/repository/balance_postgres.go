package repository

import (
	balance_api "balance-api"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) GetBalance(id int) (int, error) {
	var balance int
	query := fmt.Sprintf("SELECT balance FROM %s WHERE id = ($1)", usersTable)
	err := r.db.Get(&balance, query, id)
	return balance, err
}

func (r *BalancePostgres) UpdateBalance(id int, input balance_api.UpdateBalance, sourceOrPurpose string) error {
	var sign string
	if input.CashFlow == true {
		sign = "+"
	} else {
		sign = "-"
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET balance = balance %s $1 WHERE id = $2", usersTable, sign)
	_, err = r.db.Exec(query, input.Netto, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	transactionQuery := fmt.Sprintf("INSERT INTO %s (user_id, netto, cashflow, source_or_purpose) VALUES ($1, $2, $3, $4)", transactionsTable)
	_, err = r.db.Exec(transactionQuery, id, input.Netto, input.CashFlow, sourceOrPurpose)
	if err != nil {
		tx.Rollback()
	}

	return err
}

func (r *BalancePostgres) CreateBalance(id int, input balance_api.UpdateBalance) error {
	query := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2)", usersTable)
	_, err := r.db.Exec(query, id, input.Netto)

	return err
}
