package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Db *gorm.DB

func Connect() {
	var err error
	err = godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	dsn := os.Getenv("DSN")
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
}
