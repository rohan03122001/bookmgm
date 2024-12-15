package config

import (
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "rohan:rohan0312@tcp(127.0.0.1:3306)/social?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
