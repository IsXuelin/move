package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"move/model"
	"net/url"
)

var db *gorm.DB

func InitDB() {
	// mysql
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	var err error
	db, err = gorm.Open(driverName, args)
	//db, err := gorm.Open("mysql", "root:123456@tcp(docker.for.mac.localhost:3306)/move")
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)
	// 初始化表

	db.AutoMigrate(&model.Orders{})
	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.OrderLog{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Println("DB err = ", err)
	}
}

func SetDB(setdb *gorm.DB) {
	db = setdb
}
