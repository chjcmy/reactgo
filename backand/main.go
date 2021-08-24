package main

import (
	"context"
	api "github.com/backand/api"
	"github.com/backand/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"time"
)

type (
	user struct {
		Name string `json:"name"`
	}

	birth struct {
		Birthdays time.Time `json:"birthdays"`
	}
)

func createUser(c echo.Context) error {
	client := db.Config()
	users := &user{}
	c.Bind(users)
	ctx := context.Background()
	u, err := client.User.Create().
		SetName(users.Name).
		SetAge(time.Now()).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println("user was created: ", u)
	return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST("/users", createUser)
	e.POST("/remake", api.Remake)
	e.GET("/hosting", api.Hosting)
	e.POST("/unittest", api.UnitCreate)
	e.GET("/unitshosting", api.UnitHosting)
	e.POST("/bookcreate", api.BookCreate)
	e.GET("/bookread/:id", api.BookRead)
	e.GET("/bookshow/:num", api.BookShow)
	e.GET("/pickunitbooks/:id/:num", api.PickUnitBook)
	e.DELETE("/bookdelete/:id", api.BookDelete)
	e.PUT("/bookupdate/:id", api.BookUpdate)
	e.GET("/newbooks", api.NewBooks)
	e.POST("/login", api.Login)

	e.Logger.Fatal(e.Start(":8000"))
}
