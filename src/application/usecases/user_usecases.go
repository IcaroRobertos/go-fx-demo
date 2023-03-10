package usecases

import (
	"github.com/IcaroRobertos/go-fx-demo/src/application/interfaces"
	"github.com/IcaroRobertos/go-fx-demo/src/domain/entities"
)

type UserUseCases struct {
	UserDatabaseRepository interfaces.IUserDatabaseRepository
}

func (u *UserUseCases) Create(user entities.User) (entities.User, error) {
	return u.UserDatabaseRepository.Create(user)
}
