package handler

import (
	balance_api "balance-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) updatePurchace(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input balance_api.Purchase
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if price, ok := balance_api.Services[input.Service]; ok == true {

		var nettoPurchase balance_api.UpdateBalance = balance_api.UpdateBalance{Netto: price, CashFlow: false}

		if err := h.services.UpdateBalance(id, nettoPurchase, input.Service); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		if err := h.services.UpdatePurchase(id, input.Service); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, statusResponse{"ok"})
	} else {
		newErrorResponse(c, http.StatusBadRequest, "no such services")
		return
	}
}
