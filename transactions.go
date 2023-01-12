package balance_api

type Transactions struct {
	Id                int    `json:"id" db:"id"`
	User_id           int    `json:"user_id" db:"user_id"`
	Netto             int    `json:"netto" db:"netto"`
	Cashflow          bool   `json:"cashflow" db:"cashflow"`
	Source_or_purpose string `json:"source_or_purpose" db:"source_or_purpose"`
	TransTime         string `json:"transtime" db:"transtime"`
}
