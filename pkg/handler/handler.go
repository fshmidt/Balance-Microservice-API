package handler

import (
	"balance-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		balance := api.Group("/balance")
		{
			balance.GET("/:id", h.getBalance)
			balance.PUT("/:id", h.updateBalance)
		}
		usdBalance := api.Group("/usdBalance")
		{
			usdBalance.GET("/:id", h.getBalanceUSD)
		}
		purchase := api.Group("/purchase")
		{
			purchase.PUT("/:id", h.updatePurchace)
		}

		sendMoney := api.Group("send")
		{
			sendMoney.PUT("/:id", h.sendMoney)
		}
		transactions := api.Group("/transactions")
		{
			transactions.GET("/:id", h.transactions)
		}
		transactionsBySumm := api.Group("/trans_by_summ")
		{
			transactionsBySumm.GET("/:id", h.transactionsBySumm)
		}
	}

	return router
}
