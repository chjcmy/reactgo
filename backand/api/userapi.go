package api

import (
	"context"
	db "github.com/backand/db"
	"github.com/backand/ent/user"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type (
	info struct {
		Name     string    `json:"name"`
		Password string    `json:"password"`
		Age      time.Time `json:"age"`
		Hobby    string    `json:"hobby"`
		Lang     string    `json:"lang"`
		Gitlab   string    `json:"gitlab"`
		Github   string    `json:"github"`
	}

	host []struct {
		Name   string    `json:"name"`
		Age    time.Time `json:"age"`
		Hobby  string    `json:"hobby"`
		Lang   string    `json:"lang"`
		Gitlab string    `json:"gitlab"`
		Github string    `json:"github"`
	}
)

func Remake(c echo.Context) error {
	client := db.Config()
	infos := &info{}
	c.Bind(infos)
	ctx := context.Background()
	u, err := client.User.Create().
		SetName(infos.Name).
		SetPassword(infos.Password).
		SetAge(infos.Age).
		SetHobby(infos.Hobby).
		SetLang(infos.Lang).
		SetGithub(infos.Github).
		SetGitlab(infos.Gitlab).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println("user was created: ", u)
	return c.JSON(http.StatusOK, u)

}

func Hostinfo(c echo.Context) error {
	client := db.Config()
	ctx := context.Background()
	u, err := client.User.Query().Select(user.FieldName, user.FieldAge, user.FieldHobby, user.FieldLang, user.FieldGithub, user.FieldGitlab).Where(user.Name("최성현")).All(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println("user was created: ", u)
	return c.JSON(http.StatusOK, u)
}
