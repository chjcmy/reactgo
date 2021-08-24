package api

import (
	"context"
	"fmt"
	"github.com/backand/db"
	"github.com/backand/ent"
	"github.com/backand/ent/book"
	"github.com/backand/ent/cipher"
	"github.com/backand/ent/unit"
	"github.com/backand/ent/user"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type (
	bookmaker struct {
		Unit    int    `json:"unit" `
		Title   string `json:"title"`
		Subject string `json:"subject"`
		ID      int    `header:"id"`
	}
	bookplate struct {
		Title   string `json:"title"`
		Subject string `json:"subject"`
	}
)

func BookCreate(c echo.Context) error {
	var cipherText []struct {
		value []byte `json:"googlenum"`
	}
	var id []byte
	ctx := context.Background()
	client := db.Config()
	client.Cipher.Query().
		Where(cipher.ID(1)).
		Select(cipher.FieldDefault).
		Scan(ctx, &cipherText)

	cipherText[0].value = decrypt(cipherText[0].value, "chltjdgus123!")
	fmt.Println(cipherText[0].value)

	//booking := &bookmaker{}
	req := c.Request()
	headers := req.Header
	id = []byte(headers.Get("id"))
	fmt.Println(id)

	decode := decrypt(cipherText[0].value, "chltjdgus123!")

	fmt.Println(decode)
	//c.Bind(booking)
	//booking.ID, _ = strconv.Atoi(id)
	//fmt.Println(booking.Title, booking.Subject, booking.ID, booking.Unit)
	//cr, err := client.Book.Create().
	//	SetTitle(booking.Title).
	//	SetUnitidID(booking.Unit).
	//	//SetUseridID(func(q *ent.UserQuery) {
	//	//	q.Select(user.FieldID).Where(user.Googlenum(decode))
	//	//}).
	//	SetSubject(booking.Subject).
	//	Save(ctx)
	//if err != nil {
	//	log.Println(err)
	//	return c.JSON(http.StatusBadRequest, err)
	//}
	return c.JSON(http.StatusOK, decode)
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
		WithUnitid(func(q *ent.UnitQuery) {
			q.Select(unit.FieldContentName)
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
		WithUnitid(func(q *ent.UnitQuery) {
			q.Select(unit.FieldContentName)
		}).
		Limit(10).
		Offset(num).
		Order(ent.Desc(book.FieldID)).
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

func PickUnitBook(c echo.Context) error {
	client := db.Config()
	id, _ := strconv.Atoi(c.Param("id"))
	num, _ := strconv.Atoi(c.Param("num"))
	ctx := context.Background()
	fmt.Println(id)
	r, err := client.Book.Query().
		Select(book.FieldID, book.FieldUpdatedAt, book.FieldCreateAt, book.FieldTitle).
		Where(
			book.HasUnitidWith(
				unit.ID(id),
			)).
		WithUnitid(func(q *ent.UnitQuery) {
			q.Select(unit.FieldID, unit.FieldContentName)
		}).
		WithUserid(func(q *ent.UserQuery) {
			q.Select(user.FieldName)
		}).
		Limit(10).
		Offset(num).
		Order(ent.Desc(book.FieldID)).
		All(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, r)
}

func NewBooks(c echo.Context) error {
	client := db.Config()
	ctx := context.Background()
	r, err := client.Book.Query().
		Limit(5).
		WithUnitid(func(q *ent.UnitQuery) {
			q.Select(unit.FieldContentName)
		}).
		Order(ent.Desc(book.FieldID)).
		All(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, r)
}
