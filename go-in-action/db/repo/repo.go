package repo

import (
	"db/database"
)

type User struct {
	ID       string `gorm:"primary_key,type:varchar(50)"`
	Name     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(255)"`
	IsAdmin  bool   `gorm:"type:bit"`
}

var DB *database.Orm

func init() {
	initRepository()
	DB.Database.AutoMigrate(
		&User{},
	)
	initData()
}

func initRepository() {
	DB = &database.Orm{
		Database: database.NewDatabase(&database.DatabaseConfig{
			Host:     "localhost",
			Port:     3306,
			DBName:   "test",
			User:     "dp",
			Password: "123456",
		}),
	}
}

func initData() {
	if e := DB.First(&User{}, database.FirstOption{
		Condition: database.Condition("name = admin and is_admin = 1"),
	}); e != nil {
		DB.Create(&User{
			Name:     "admin",
			Password: "123456",
			IsAdmin:  true,
		})
	}
}
