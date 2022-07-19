package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mskydream/qr-code/cmd/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
	}
	return router
}
