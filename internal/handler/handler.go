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

//	type Segment interface {
//		CreateSegment(segment.Segment) (segment.Segment, error)
//		DeleteSegment(string) error
//	}
//
//	type User interface {
//		AddUser(segment.User) error
//		CheckUser(segment.User) ([]segment.Segment, error)
//	}
//
//	type ActiveSegment interface {
//		AddActiveSegment(segment.ActiveSegment) (segment.ActiveSegment, error)
//		DeleteActiveSegment(segment.ActiveSegment) error
//	}
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
