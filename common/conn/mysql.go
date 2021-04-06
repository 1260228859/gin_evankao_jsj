package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func InitMysqlClient() (db *gorm.DB, err error) {
	db, err = gorm.Open(viper.GetString("datasource.driverName"), viper.GetString("datasource.uri"))
	return db, err
}
