package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var db *gorm.DB

func Connect() {
	dsn := os.Getenv("DSN")
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
}
