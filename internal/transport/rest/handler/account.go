package handler

import (
	"coins-app/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required"`
}

type createAccountResponse struct {
	Id        int    `json:"id"`
	Currency  string `json:"currency"`
	Balance   int64  `json:"balance"`
	CreatedAt string `json:"created_at"`
}

// @Summary Create Account
// @Security ApiKeyAuth
// @Description Create Account
// @Tags account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body createAccountRequest true "Account"
// @Success 200 {object} createAccountResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/accounts [post]
func (h *Handler) createAccount(c *gin.Context) {
	var request createAccountRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	arg := core.Account{
		UserId:   userId,
		Currency: request.Currency,
		Balance:  0,
	}

	id, err := h.services.Account.CreateAccount(arg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, createAccountResponse{
		Id:        id,
		Currency:  arg.Currency,
		Balance:   arg.Balance,
		CreatedAt: time.Time{}.String(),
	})
}

type getAccountByIdRequest struct {
	Id int `uri:"id" binding:"required,min=1"`
}

func (h *Handler) getAccountById(c *gin.Context) {
	var request getAccountByIdRequest

	if err := c.BindUri(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	account, err := h.services.Account.GetAccountById(request.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if account.UserId != userId {
		newErrorResponse(c, http.StatusForbidden, "not your account")
		return
	}

	c.JSON(http.StatusOK, account)
}

type getAccountsResponse struct {
	Accounts []core.Account `json:"accounts"`
}

// @Summary Get Accounts
// @Security ApiKeyAuth
// @Description Get Accounts
// @Tags account
// @ID get-accounts
// @Accept json
// @Produce json
// @Success 200 {object} getAccountsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/accounts [get]
func (h *Handler) getAccounts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	accounts, err := h.services.Account.GetAccounts(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAccountsResponse{
		Accounts: accounts,
	})
}

type updateAccountRequest struct {
	Balance  int64  `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

// @Summary Update Account
// @Security ApiKeyAuth
// @Description Update Account
// @Tags account
// @ID update-account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param input body updateAccountRequest true "Account"
// @Success 200 {object} core.Account
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/accounts/{id} [put]
func (h *Handler) updateAccount(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid account id param")
		return
	}

	var request updateAccountRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	account, err := h.services.Account.GetAccountById(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if account.UserId != userId {
		newErrorResponse(c, http.StatusForbidden, "not your account")
		return
	}

	if account.Currency != request.Currency {
		newErrorResponse(c, http.StatusBadRequest, "currency mismatch")
		return
	}

	if request.Balance < 0 {
		newErrorResponse(c, http.StatusBadRequest, "balance cannot be negative")
		return
	}

	arg := core.Account{
		Id:       accountId,
		Balance:  request.Balance,
		Currency: request.Currency,
	}

	err = h.services.Account.UpdateAccount(arg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, arg)
}
