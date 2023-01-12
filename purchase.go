package balance_api

type Purchase struct {
	Service string `json:"services" db:"services" binding:"required"`
}
