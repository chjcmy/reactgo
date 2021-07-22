package main

import (
	"context"
	"fmt"
	userapi "github.com/backand/api"
	"github.com/backand/db"
	"github.com/backand/ent"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"math/rand"
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

func BirthDay(c echo.Context) error {
	births := &birth{}
	c.Bind(births)
	log.Println("user was created: ", births)

	year := diff(births.Birthdays, time.Now())

	fmt.Printf("You are %d years old.",
		year)
	return c.JSON(http.StatusOK, year)
}

func diff(a, b time.Time) (year int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, _, _ := a.Date()
	y2, _, _ := b.Date()

	year = int(y2 - y1) //nolint:unconvert

	return
}

func main() {
	e := echo.New()

	client, err := ent.Open("mysql", "cshcmi:chltjdgus123!@tcp(choi1994.iptime.org:1994)/blog?charset=utf8mb4&parseTime=True&loc=Local")
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
	e.POST("/Birthday", BirthDay)
	e.POST("/remake", userapi.Remake)
	e.GET("/hostinfo", userapi.Hostinfo)

	e.GET("/jsonp", func(c echo.Context) error {
		callback := c.QueryParam("callback")
		var content struct {
			Response  string    `json:"response"`
			Timestamp time.Time `json:"timestamp"`
			Random    int       `json:"random"`
		}
		content.Response = "Sent via JSONP"
		content.Timestamp = time.Now().UTC()
		content.Random = rand.Intn(1000) //nolint:gosec
		return c.JSONP(http.StatusOK, callback, &content)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
