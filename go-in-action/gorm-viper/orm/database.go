package orm

import (
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

type DatabaseConfig struct {
	Type     DatabaseType `json:"type"`
	Host     string       `json:"host"`
	Port     int          `json:"port"`
	Database string       `json:"database"`
	Username string       `json:"username"`
	Password string       `json:"password"`
}

type DatabaseType string

const (
	Mysql      DatabaseType = "mysql"
	Sqlserver  DatabaseType = "sqlserver"
	Postgresql DatabaseType = "postgresql"
	Sqlite     DatabaseType = "sqlite"
)

func NewDatabase(dbConf DatabaseConfig) *Database {
	var db *Database
	switch dbConf.Type {
	case Mysql:
		db = newMysql(dbConf)
	case Postgresql:
		db = newPgsql(dbConf)
	case Sqlserver:
		db = newSqlserver(dbConf)
	case Sqlite:
		panic("unimplemented.")
	default:
		panic("unsupported database.")
	}
	return db
}

