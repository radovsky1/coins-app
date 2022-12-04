package handler

import (
	"coins-app/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signUpResponse struct {
	Id int `json:"id"`
}

// @Summary Sign Up
// @Tags auth
// @Description Create new user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param user body core.User true "User"
// @Success 200 {object} signUpResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var request core.User

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := h.services.Authorization.CreateUser(request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, signUpResponse{
		Id: id,
	})
}

type signInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signInResponse struct {
	Token string `json:"token"`
}

// @Summary Sign In
// @Tags auth
// @Description Sign in user
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param user body signInRequest true "User"
// @Success 200 {object} signInResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var request signInRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(request.Username, request.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, signInResponse{
		Token: token,
	})
}
