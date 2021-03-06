package api

import (
	"context"
	"fmt"
	"github.com/backand/db"
	"github.com/backand/ent"
	"github.com/backand/ent/unit"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type AutoGenerated struct {
	Units []struct {
		Contents     string `json:"contents"`
		ContentNames string `json:"content_names"`
	} `json:"units"`
}

func UnitCreate(c echo.Context) error {
	client := db.Config()
	u := &AutoGenerated{Units: nil}
	ctx := context.Background()
	if err := c.Bind(u); err != nil { // here unmarshal request body into p
		return c.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Printf("%v", u)
	unit := make([]*ent.UnitCreate, len(u.Units))
	for i := 0; i < len(u.Units); i++ {
		unit[i] = client.Unit.Create().SetContent(u.Units[i].Contents).SetContentName(u.Units[i].ContentNames)
	}
	units, err := client.Unit.CreateBulk(unit...).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, units)
}

func UnitHosting(c echo.Context) error {
	client := db.Config()
	ctx := context.Background()
	units, err := client.Unit.Query().Select(unit.FieldID, unit.FieldContent, unit.FieldContentName).All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, units)
}
