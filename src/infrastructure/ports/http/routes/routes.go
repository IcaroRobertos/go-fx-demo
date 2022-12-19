package routes

import (
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/ports/http/controllers"
	"github.com/gin-gonic/gin"
)

type RouteControllers struct {
	RootController *controllers.RootController
	UserController *controllers.UserController
}

func NewRouteControllers(rootController *controllers.RootController, userController *controllers.UserController) *RouteControllers {
	return &RouteControllers{
		RootController: rootController,
		UserController: userController,
	}
}

func Router(r *gin.Engine, controllers *RouteControllers) {
	r.GET("/", controllers.RootController.Root)

	r.POST("/user", controllers.UserController.Create)
}
