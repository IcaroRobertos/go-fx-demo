package main

import (
	"github.com/IcaroRobertos/go-fx-demo/src/application/usecases"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/io/databases"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/io/webservers"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/ports/http/controllers"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/ports/http/routes"
	"github.com/IcaroRobertos/go-fx-demo/src/infrastructure/repositories"
)

func main() {
	mainDatabase := databases.GormConnect()

	controllers := routes.RouteControllers{
		RootController: controllers.RootController{},
		UserController: controllers.UserController{
			UserUseCases: usecases.UserUseCases{
				UserDatabaseRepository: &repositories.UserDatabaseRepository{
					MainDatabase: mainDatabase,
				},
			},
		},
	}

	webservers.NewGinWebServer(controllers)
}
