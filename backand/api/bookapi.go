package api

import (
	"context"
	"github.com/backand/db"
	"github.com/backand/ent"
	"github.com/backand/ent/book"
	"github.com/backand/ent/user"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type (
	bookmaker struct {
		Title   string `json:"title"`
		Subject string `json:"subject"`
		Units   int    `json:"units"`
		User    int    `json:"user"`
	}
	bookplate struct {
		Title   string `json:"title"`
		Subject string `json:"subject"`
	}
)

func BookCreate(c echo.Context) error {
	client := db.Config()
	booking := &bookmaker{}
	c.Bind(booking)
	ctx := context.Background()
	cr, err := client.Book.Create().
		SetTitle(booking.Title).
		SetUnitidID(booking.Units).
		SetUseridID(booking.User).
		SetSubject(booking.Subject).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println("user was created: ", cr)
	return c.JSON(http.StatusOK, cr)
}

func BookRead(c echo.Context) error {
	client := db.Config()
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := context.Background()
	r, err := client.Book.Query().
		Where(book.ID(id)).
		WithUserid(func(q *ent.UserQuery) {
			q.Select(user.FieldName)
		}).
		Only(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, r)
}

func BookShow(c echo.Context) error {
	client := db.Config()
	num, _ := strconv.Atoi(c.Param("num"))
	ctx := context.Background()
	r, err := client.Book.Query().
		WithUserid(func(q *ent.UserQuery) {
			q.Select(user.FieldName)
		}).
		Limit(10).
		Offset(num).
		All(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, r)
}

func BookUpdate(c echo.Context) error {
	client := db.Config()
	id, _ := strconv.Atoi(c.Param("id"))
	booking := &bookplate{}
	c.Bind(booking)
	ctx := context.Background()
	u, err := client.Book.
		UpdateOneID(id).
		SetTitle(booking.Title).
		SetSubject(booking.Subject).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, u)
}

func BookDelete(c echo.Context) error {
	client := db.Config()
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))
	err := client.Book.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "success")
}