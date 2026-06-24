package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoalHandler struct{}

func NewGoalHandler() *GoalHandler { return &GoalHandler{} }

func (h *GoalHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"goals": []interface{}{}})
}

func (h *GoalHandler) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "goal created"})
}

func (h *GoalHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"goal": nil})
}

func (h *GoalHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "goal updated"})
}

func (h *GoalHandler) Delete(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
