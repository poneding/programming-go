package entities

import "gorm-viper/orm"

type User struct {
	orm.Model
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Age  int    `json:"age"`
}
