package repositories

import (
	"errors"

	"github.com/IcaroRobertos/go-fx-demo/src/domain/entities"
	"gorm.io/gorm"
)

type UserDatabaseRepository struct {
	MainDatabase *gorm.DB
}

func (udr *UserDatabaseRepository) Create(user entities.User) (entities.User, error) {
	result := udr.MainDatabase.Create(&user)
	if result.Error != nil {
		return entities.User{}, errors.New(result.Error.Error())
	}

	return user, nil
}
