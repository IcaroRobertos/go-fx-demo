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

	rootController := controllers.RootController{}
	userController := controllers.UserController{
		UserUseCases: usecases.UserUseCases{
			UserDatabaseRepository: &repositories.UserDatabaseRepository{
				MainDatabase: mainDatabase,
			},
		},
	}

	controllers := routes.RouteControllers{
		RootController: rootController,
		UserController: userController,
	}

	webservers.NewGinWebServer(controllers)
}
