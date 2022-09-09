package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mskydream/ashyq/docs"
	"github.com/mskydream/ashyq/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	docs.SwaggerInfo.BasePath = "/api"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
			realEstate.POST("", h.createRealEstate)
			realEstate.PUT("/:id", h.updateRealEstate)
			realEstate.DELETE("/:id", h.deleteRealEstate)
			realEstate.StaticFS("/qr_code", http.Dir("./cmd/qr"))
		}

		visit := api.Group("/visit")
		{
			visit.GET("/:id", h.getVisit)
			visit.GET("", h.getAllVisits)
			visit.Group("/qr_code").GET("/:id", h.getVisitByQrCode)
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
