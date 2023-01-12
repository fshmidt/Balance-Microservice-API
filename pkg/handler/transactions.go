package handler

import (
	balance_api "balance-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	BYDATE = 0
	BYSUMM = 1
)

type getAllTransResponse struct {
	Data []balance_api.Transactions `json:"transactions"`
}

func (h *Handler) transactions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	lists, err := h.services.GetTransactions(id, BYDATE)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllTransResponse{
		Data: lists,
	})
}

func (h *Handler) transactionsBySumm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	lists, err := h.services.GetTransactions(id, BYSUMM)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllTransResponse{
		Data: lists,
	})
}
