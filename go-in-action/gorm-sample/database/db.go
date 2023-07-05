package database

import (
	"gorm.io/gorm"
)

type database struct {
	*gorm.DB
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbName"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func NewDatabase(dbConf *DatabaseConfig) *database {
	db := newMysql(dbConf)
	return db
}
