package handler

import (
	service "avito-sest-segments/internal/service"

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

	user := router.Group("/user")
	{
		user.GET("/", h.AddUser)
		user.POST("/", h.CheckUser)
	}
	product := router.Group("/order")
	{
		product.POST("/reserve", h.CreateOrder)
		product.POST("/revenue", h.RevenueOrder)
	}

	return router
}
