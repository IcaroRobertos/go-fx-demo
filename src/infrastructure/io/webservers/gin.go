package webservers

import (
	"github.com/gin-gonic/gin"
)

func NewGinWebServer() *gin.Engine {
	r := gin.Default()

	return r
}
