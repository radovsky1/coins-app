package handler

import (
	"coins-app/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createTransferRequest struct {
	FromAccountID int    `json:"from_account_id" binding:"required"`
	ToAccountID   int    `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required"`
}

type createTransferResponse struct {
	Id int `json:"id"`
}

// @Summary Create Transfer
// @Security ApiKeyAuth
// @Description Create Transfer
// @Tags transfer
// @ID create-transfer
// @Accept json
// @Produce json
// @Param input body createTransferRequest true "Transfer"
// @Success 200 {object} createTransferResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/transfers [post]
func (h *Handler) createTransfer(c *gin.Context) {
	var request createTransferRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	arg := core.Transfer{
		FromAccountID: request.FromAccountID,
		ToAccountID:   request.ToAccountID,
		Amount:        request.Amount,
		Currency:      request.Currency,
	}

	transferId, err := h.services.Transfer.CreateTransfer(arg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, createTransferResponse{
		Id: transferId,
	})
}

type getTransferByIdRequest struct {
	Id int `uri:"id" binding:"required,min=1"`
}

func (h *Handler) getTransferById(c *gin.Context) {
	var request getTransferByIdRequest

	if err := c.ShouldBindUri(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transfer, err := h.services.Transfer.GetTransferById(request.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transfer)
}

// @Summary Get Transfers
// @Security ApiKeyAuth
// @Description Get Transfers
// @Tags transfer
// @ID get-transfers
// @Accept json
// @Produce json
// @Success 200 {object} []core.Transfer
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/transfers [get]
func (h *Handler) getTransfers(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	transfers, err := h.services.Transfer.GetTransfers(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transfers)
}
