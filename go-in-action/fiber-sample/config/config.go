package config

import (
	"fiber-sample/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error

	Database, err = gorm.Open(sqlite.Open("fiber-sample.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = Database.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}

	return nil
}
