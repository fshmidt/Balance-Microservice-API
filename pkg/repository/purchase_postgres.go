package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PurchasePostgres struct {
	db *sqlx.DB
}

func NewPurchasePostgres(db *sqlx.DB) *PurchasePostgres {
	return &PurchasePostgres{db: db}
}

func (r *PurchasePostgres) UpdatePurchase(id int, srv string) error {
	query := fmt.Sprintf("UPDATE %s SET services = CASE WHEN services IS NULL THEN '%s' ELSE CONCAT(services,', %s') END  WHERE id = $1", usersTable, srv, srv)
	_, err := r.db.Exec(query, id)

	return err
}
