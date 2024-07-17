package main

import (
	"github.com/ipe-dev/menu-memo/backend/src/database"
	"github.com/ipe-dev/menu-memo/backend/src/database/model"
)

func main() {
	database.Connect()
	database.Db.AutoMigrate(&model.User{})
}