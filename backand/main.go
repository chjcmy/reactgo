package main

import (
	api "github.com/backand/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type (
	jwtCustomClaims struct {
		Name  string `json:"name"`
		jwt.StandardClaims
	}
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))


	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/remake", api.Remake)
	e.GET("/hosting", api.Hosting)
	e.POST("/unittest", api.UnitCreate)
	e.GET("/unitshosting", api.UnitHosting)
	e.POST("/bookcreate", api.BookCreate)
	e.GET("/bookread/:id", api.BookRead)
	e.GET("/bookshow", api.BookShow)
	e.GET("/pickunitbooks/:id/:num", api.PickUnitBook)
	e.DELETE("/bookdelete/:id", api.BookDelete)
	e.PUT("/bookupdate/:id", api.BookUpdate)
	e.GET("/newbooks", api.NewBooks)
	e.POST("/login", api.Login)
	r := e.Group("/relogin")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		SigningKey: []byte("chltjdgus123!"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", api.ReLogin)



	e.Logger.Fatal(e.Start(":8000"))
}
