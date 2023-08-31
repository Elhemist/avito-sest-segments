package handler

import (
	segment "avito-sest-segments/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
func (h *Handler) CheckUser(c *gin.Context) {
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
	c.JSON(http.StatusOK, user)

}
func (h *Handler) AddSegments(c *gin.Context) {
	var input segment.SegmentsToUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input.Delete)
	err := h.services.User.AddSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Done",
	})
}
