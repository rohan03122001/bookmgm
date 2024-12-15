package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
    // db holds the database connection instance
    db *gorm.DB
)

// Connect establishes a connection to MySQL database using GORM
func Connect() {
    // Establish database connection with connection string
    // Format: username:password@tcp(host:port)/dbname?params
    d, err := gorm.Open("mysql", "rohan:rohan0312@tcp(127.0.0.1:3306)/social?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    db = d
}

// GetDB returns the active database connection
func GetDB() *gorm.DB {
    return db
}