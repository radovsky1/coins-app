package handler

import (
	"coins-app/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getCoinPriceRequest struct {
	Coin string `json:"coin" binding:"required"`
}

type getCoinPriceResponse struct {
	Prices []core.SymbolPrice `json:"prices"`
}

func (h *Handler) getCoinPrice(c *gin.Context) {
	var request getCoinPriceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	coinPrices, err := h.services.Coin.GetCoinPrices(request.Coin)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error getting coin prices")
		return
	}

	c.JSON(http.StatusOK, getCoinPriceResponse{
		Prices: coinPrices,
	})
}
