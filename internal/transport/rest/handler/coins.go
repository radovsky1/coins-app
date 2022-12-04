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

// @Summary Get Coin Price
// @Security ApiKeyAuth
// @Description Get Coin Price
// @Tags coin
// @ID get-coin-price
// @Accept json
// @Produce json
// @Param input body getCoinPriceRequest true "Coin"
// @Success 200 {object} getCoinPriceResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/coin/price [get]
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
