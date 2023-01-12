package balance_api

type UpdateBalance struct {
	Netto    int  `json:"netto" db:"balance" binding:"required"`
	CashFlow bool `json:"cashflow" db:"cashflow"`
}
