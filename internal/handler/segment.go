package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Segment struct {
	Name string
}

func (h *Handler) CreateSegment(c *gin.Context) {
	var input Segment
	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Segment.CreateSegment(input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"SegmentId": id,
	})
}
func (h *Handler) DeleteSegment(c *gin.Context) {
	var input Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Segment.DeleteSegment(input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Segment deleted",
	})
}
