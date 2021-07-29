package main

import (
	"context"
	bookapi "github.com/backand/api"
	unitapi "github.com/backand/api"
	userapi "github.com/backand/api"
	"github.com/backand/db"
	"github.com/backand/ent"
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

	client, err := ent.Open("mysql", "cshcmi:chltjdgus123!@tcp(choi1994.iptime.org:1994)/blog?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Panicf("failed creating schema resources: %v", err)
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", createUser)
	e.POST("/remake", userapi.Remake)
	e.GET("/hosting", userapi.Hosting)
	e.POST("/unittest", unitapi.UnitCreate)
	e.GET("/unitshosting", unitapi.UnitHosting)
	e.POST("/bookcreate", bookapi.BookCreate)
	e.GET("/bookread/:id", bookapi.BookRead)
	e.GET("/bookshow/:num", bookapi.BookShow)
	e.DELETE("/bookdelete/:id", bookapi.BookDelete)
	e.PUT("/bookupdate/:id", bookapi.BookUpdate)

	e.Logger.Fatal(e.Start(":8000"))
}
