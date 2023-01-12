package handler

import (
	balance_api "balance-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) sendMoney(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid sender id param")
		return
	}
	var input balance_api.Send
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Netto <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "Netto has to be a positive number")
		return
	}
	if input.ReacherId <= 0 || input.ReacherId == id {
		newErrorResponse(c, http.StatusBadRequest, "Invalid reacher id param")
		return
	}

	var nettoSender balance_api.UpdateBalance = balance_api.UpdateBalance{Netto: input.Netto, CashFlow: false}
	var nettoReacher balance_api.UpdateBalance = balance_api.UpdateBalance{Netto: input.Netto, CashFlow: true}
	if err := h.services.UpdateBalance(id, nettoSender, strconv.Itoa(input.ReacherId)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.UpdateBalance(input.ReacherId, nettoReacher, c.Param("id")); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
