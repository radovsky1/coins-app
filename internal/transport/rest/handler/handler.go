package handler

import (
	"coins-app/internal/service"
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		accounts := api.Group("/accounts")
		{
			accounts.POST("/", h.createAccount)
			accounts.GET("/", h.getAccounts)
			accounts.GET("/:id", h.getAccountById)
			accounts.PUT("/:id", h.updateAccount)
		}

		transfers := api.Group("/transfers")
		{
			transfers.POST("/", h.createTransfer)
			transfers.GET("/", h.getTransfers)
			transfers.GET("/:id", h.getTransferById)
		}

		coins := api.Group("/coin")
		{
			coins.GET("/price", h.getCoinPrice)
		}
	}

	return router
}
