package balance_api

type Send struct {
	Netto     int `json:"netto" binding:"required"`
	ReacherId int `json:"reacherid" binding:"required"`
}
