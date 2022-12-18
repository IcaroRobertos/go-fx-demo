package databases

import (
	"github.com/IcaroRobertos/go-fx-demo/src/domain/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GormConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.User{})

	return db
}
