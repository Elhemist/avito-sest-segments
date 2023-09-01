package handler

import (
	service "avito-sest-segments/internal/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Elhemist/avito-sest-segments/docs"
)

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
