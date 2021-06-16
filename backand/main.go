package main

import (
	api3 "backand/api"
	"github.com/gin-gonic/gin"
)

func main() {
	//db := db3.Connect()
	//
	//db.AutoMigrate(&db3.User{}, &db3.Book{}, &db3.BookSubject{})

	//db.Create(&db3.Book{
	//	Model:       gorm.Model{},
	//	Title:       "math",
	//	Content:     "mula",
	//	BookName:    1,
	//	User:        db3.User{},
	//	BookCode:    3,
	//	BookSubject: db3.BookSubject{},
	//})

	r := gin.Default()

	api := r.Group("/api")

	api.Use()
	{
		user := api.Group("/user")
		user.GET("/userinfo", api3.Userinfo)
	}

	api.Use()
	{
		subjects := api.Group("/book")
		subjects.GET("/:id", api3.ShowBook)
		subjects.GET("/books", api3.ShowTitles)
	}

	api.Use()
	{
		subjects := api.Group("/subjects")
		subjects.GET("/subject", api3.ShowSubject)
	}

	r.Run("0.0.0.0:8000")
}
