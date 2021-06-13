package main

import (
	api3 "backand/api"
	db3 "backand/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db3.Connect()

	db.AutoMigrate(&db3.User{}, &db3.Book{}, &db3.BookSubject{})

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/userinfo", api3.Userinfo)
	}
	r.Run("0.0.0.0:8000")
}
