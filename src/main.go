package main

import (
	"github.com/IcaroRobertos/go-fx-demo/src/application/usecases"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/io/databases"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/io/webservers"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/ports/http/controllers"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/ports/http/routes"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(databases.GormConnect),
		fx.Provide(repositories.NewUserDatabaseRepository),
		fx.Provide(usecases.NewUserUseCases),
		fx.Provide(controllers.NewRootController),
		fx.Provide(controllers.NewUserController),
		fx.Provide(routes.NewRouteControllers),
		fx.Provide(webservers.NewGinWebServer),
		fx.Invoke(routes.Router),
		fx.Invoke(func(r *gin.Engine) {
			r.Run()
		}),
	).Run()
}
