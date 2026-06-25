package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler { return &AuthHandler{} }

// Login is a stub — full JWT issuance to be implemented.
func (h *AuthHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": "stub-jwt-token"})
}
