package handler

import (
	segment "avito-sest-segments/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type resp struct {
	Res string `json:"message"`
}
type segResp struct {
	SegmentNames []string `json:"userSegmentNames"`
}

// @Summary AddUser
// @Tags user
// @Description Add user to the base
// @ID add-user
// @Accept  json
// @Produce  json
// @Param input body segment.User true "User id"
// @Success 200 {object} resp
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/ [post]
func (h *Handler) AddUser(c *gin.Context) {
	var input segment.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.User.AddUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	////////////////////////////////!!!!!!!!!!!!!!!!!!!
	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}

// @Summary CheckUser
// @Tags user
// @Description Get all segments for user
// @ID check-user
// @Accept  json
// @Produce  json
// @Param input body segment.User true "User id"
// @Success 200 {object} segResp
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/ [get]
func (h *Handler) CheckUser(c *gin.Context) {
	var input segment.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	segArr, err := h.services.User.CheckUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"userSegmentNames": segArr,
	})

}

// @Summary AddSegments
// @Tags user
// @Description Add segments to user
// @ID add-segment
// @Accept  json
// @Produce  json
// @Param input body segment.SegmentsToUpdate true "Segments info"
// @Success 200 {object} resp
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/ [patch]
func (h *Handler) AddSegments(c *gin.Context) {
	var input segment.SegmentsToUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.User.AddSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Done",
	})
}
