package main

import (
	api3 "backand/api"
	db3 "backand/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db3.Connect()

	db.AutoMigrate( &db3.User{}, &db3.Book{}, &db3.BookSubject{})

	//db.Create(&db3.User{
	//	Model:   gorm.Model{},
	//	UserName:    "최성현",
	//	Email:   "chjcmy@gmail.com",
	//	Github:  "https://github.com/chjcmy",
	//	Gitlab:  "https://gitlab.com/chjcmy",
	//	Manager: 0,
	//})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")
	{
		api.POST("/userinfo", api3.Userinfo)
	}
	r.Run("0.0.0.0:8000")
}
