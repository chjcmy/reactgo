package main

import (
db "backand/db"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Code  string
	Price uint
}

func main() {
	db := db.Connect()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		db.AutoMigrate(&Product{})

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8000")
}