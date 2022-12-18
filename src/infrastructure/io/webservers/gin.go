package webservers

import (
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/ports/http/routes"
	"github.com/gin-gonic/gin"
)

func NewGinWebServer(controllers routes.RouteControllers) {
	r := gin.Default()

	routes.Router(r, controllers)

	r.Run()
}
