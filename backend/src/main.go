package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipe-dev/menu-memo/backend/src/database"
)

func init() {
	database.Connect()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
