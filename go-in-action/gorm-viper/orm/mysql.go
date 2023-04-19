package orm

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func newMysql(dbConf DatabaseConfig) *Database {
	if dbConf == (DatabaseConfig{}) {
		panic("Read mysql database config failed.")
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&timeout=60m",
		dbConf.Username,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database)

	var logLevel = logger.Error
	if os.Getenv("APP_ENVIRONMENT") == "Debugging" {
		logLevel = logger.Info
	}

	mysqlConf := mysql.Config{
		DSN:                      conn,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true,
	}

	gormConf := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logLevel),
	}

	db, _ := gorm.Open(mysql.New(mysqlConf), &gormConf)

	if db != nil {
		pool, _ := db.DB()
		pool.SetMaxIdleConns(10)
		pool.SetMaxOpenConns(100)
		pool.SetConnMaxLifetime(time.Hour)
	}
	return &Database{db}
}
