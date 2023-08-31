package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := router.Group("/user")
	{
		user.GET("/", h.AddUser())
		user.POST("/", h.UpdateBalance)
	}
	product := router.Group("/order")
	{
		product.POST("/reserve", h.CreateOrder)
		product.POST("/revenue", h.RevenueOrder)
	}

	return router
}
