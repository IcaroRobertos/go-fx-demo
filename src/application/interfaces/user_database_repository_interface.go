package interfaces

import "github.com/IcaroRobertos/go-fx-demo/src/domain/entities"

type IUserDatabaseRepository interface {
	Create(entities.User) (entities.User, error)
}
