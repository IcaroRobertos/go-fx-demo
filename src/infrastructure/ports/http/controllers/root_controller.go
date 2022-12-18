package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RootController struct{}

func (rt *RootController) Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
