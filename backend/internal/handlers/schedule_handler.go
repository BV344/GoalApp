package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct{}

func NewScheduleHandler() *ScheduleHandler { return &ScheduleHandler{} }

func (h *ScheduleHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"slots": []interface{}{}})
}

func (h *ScheduleHandler) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "slot created"})
}

func (h *ScheduleHandler) Delete(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
