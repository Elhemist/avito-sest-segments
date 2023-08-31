package handler

import (
	segment "avito-sest-segments/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
)

func (h *Handler) AddUser(c *gin.Context) {
	var input segment.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.User.CheckUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"userId": user.Id,
	})
}
func (h *Handler) CheckUser(c *gin.Context) {
	var input []segment.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.User.CheckUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)

}
