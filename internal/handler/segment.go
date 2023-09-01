package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Segment struct {
	Name string `json:"segName"`
}
type SegmentId struct {
	Id int `json:"SegmentId"`
}

// @Summary CreateSegment
// @Tags segment
// @Description Add new segment to base
// @ID check-order
// @Accept  json
// @Produce  json
// @Param input body Segment true "Segment name"
// @Success 200 {object} atest.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment/ [post]
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

// @Summary DeleteSegment
// @Tags segment
// @Description Delete segment from the base
// @ID check-order
// @Accept  json
// @Produce  json
// @Param input body Segment true "Segment name"
// @Success 200 {string} "Segment deleted"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment/ [delete]
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
