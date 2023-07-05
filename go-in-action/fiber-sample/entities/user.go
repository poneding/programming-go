package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)" json:"name"`
	Age  int    `gorm:"type:int" json:"age"`
}
