package api

import (
	"context"
	db "github.com/backand/db"
	"github.com/backand/ent/user"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
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

	googleNum struct {
		Num string `json:"num"`
	}
)

func Remake(c echo.Context) error {
	client := db.Config()
	infos := &info{}
	c.Bind(infos)
	ctx := context.Background()
	u, err := client.User.Create().
		SetName(infos.Name).
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

type jwtCustomClaims  struct {
	Name  string `json:"name"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	var userName []struct {
		GoogleNum string `json:"googlenum"`
	}
	client := db.Config()
	nums := &googleNum{}
	c.Bind(nums)
	ctx := context.Background()
	err := client.User.Query().
		Where(user.Googlenum(nums.Num)).
		Select(user.FieldGooglenum).
		Scan(ctx, &userName)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	var claims = &jwtCustomClaims{
		userName[0].GoogleNum,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("chltjdgus123!"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func ReLogin(c echo.Context) error {
	var userName  []struct {
		GoogleNum string `json:"googlenum"`
	}
	googlenums := c.Get("user").(*jwt.Token)
	claims := googlenums.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	client := db.Config()
	ctx := context.Background()
	err := client.User.Query().
		Where(user.Googlenum(name)).
		Select(user.FieldGooglenum).
		Scan(ctx, &userName)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": googlenums.Raw,
	})
}

func Hosting(c echo.Context) error {
	client := db.Config()
	ctx := context.Background()
	var host []struct {
		Name   string `json:"name"`
		Age    string `json:"age"`
		Hobby  string `json:"hobby"`
		Lang   string `json:"lang"`
		Github string `json:"github"`
		Gitlab string `json:"gitlab"`
		Email  string `json:"email"`
	}
	err := client.User.Query().
		Where(user.Name("최성현")).
		Select(
			user.FieldName,
			user.FieldAge,
			user.FieldHobby,
			user.FieldLang,
			user.FieldGithub,
			user.FieldGitlab,
			user.FieldEmail,
		).
		Scan(ctx, &host)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, host)
	}

	host[0].Age = diff(host[0].Age, time.Now())

	return c.JSON(http.StatusOK, host)
}

func diff(a string, b time.Time) (year string) {
	y1, _ := strconv.Atoi(a[:4])
	y2, _, _ := b.Date()
	year = strconv.Itoa(y2 - y1)

	return
}
