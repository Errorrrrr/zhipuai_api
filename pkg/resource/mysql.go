package resource

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"zhipuai_api/internal/config"
)

func openDB(c config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Mysql.UserName, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.DBName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,     // Slow SQL threshold
			LogLevel:      logger.Info,     // Log level
			Colorful:      c.Mysql.DBDebug, // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(c.Mysql.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(c.Mysql.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(c.Mysql.MaxLifetime) * time.Minute)
	return db
}

func getMysqlDB(c config.DBConfig) *gorm.DB {
	return openDB(c)
}
