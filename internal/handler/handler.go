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
		user.POST("/", h.AddUser)
		user.PATCH("/", h.AddSegments)
		user.GET("/", h.CheckUser)
	}
	segment := router.Group("/segment")
	{
		segment.POST("/", h.CreateSegment)
		segment.DELETE("/", h.DeleteSegment)
	}

	return router
}
