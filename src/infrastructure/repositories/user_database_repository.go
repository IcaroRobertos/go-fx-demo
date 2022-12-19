package repositories

import (
	"errors"

	"github.com/IcaroRobertos/go-fx-demo/src/application/interfaces"
	"github.com/IcaroRobertos/go-fx-demo/src/domain/entities"
	"gorm.io/gorm"
)

type UserDatabaseRepository struct {
	MainDatabase *gorm.DB
}

func NewUserDatabaseRepository(db *gorm.DB) interfaces.IUserDatabaseRepository {
	return &UserDatabaseRepository{
		MainDatabase: db,
	}
}

func (udr *UserDatabaseRepository) Create(user entities.User) (entities.User, error) {
	result := udr.MainDatabase.Create(&user)
	if result.Error != nil {
		return entities.User{}, errors.New(result.Error.Error())
	}

	return user, nil
}
