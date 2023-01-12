package handler

import (
	balance_api "balance-api"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	balance, err := h.services.GetBalance(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}

func (h *Handler) getBalanceUSD(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	balance, err := h.services.GetBalance(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "No such id")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"balance in USD": fmt.Sprintf("%.1fl", float64(balance)/70.0),
	})
}

func (h *Handler) updateBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input balance_api.UpdateBalance
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Netto <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "Netto has to be a positive number")
		return
	}
	if err := h.services.UpdateBalance(id, input, "developer"); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
