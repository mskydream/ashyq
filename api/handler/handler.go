package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mskydream/ashyq/api/service"
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
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		realEstate := api.Group("/real-estate")
		{
			realEstate.GET("/", h.getAllRealEstates)
			realEstate.GET("/:id", h.getRealEstate)
			realEstate.POST("/", h.createRealEstate)
			realEstate.PUT("/:id", h.updateRealEstate)
			realEstate.DELETE("/:id", h.deleteRealEstate)
			realEstate.StaticFS("/qr-code", http.Dir("./cmd/qr"))
		}

		visit := api.Group("/visit")
		{
			visit.POST("/", h.createVisit)
			visit.GET("/:id", h.getVisit)
			visit.GET("/", h.getAllVisits)
		}

	}

	// admin := router.Group("/admin")
	// {
	// 	admin.POST("sign-in", h.adminSignIn)
	// }

	// main := router.Group("/main", h.adminIdentity)
	// {
	// 	main.GET("/", h.getAllRealEstates)
	// }

	// amdin.GET("/:id", h.getUser)
	// amdin.POST("/", h.createUser)
	// amdin.PUT("/:id", h.updateUser)
	// amdin.DELETE("/:id", h.deleteUser)

	return router
}
