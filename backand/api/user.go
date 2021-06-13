package api

import (
	db3 "backand/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Userinfo(c *gin.Context) {

	var user db3.User
		db := db3.Connect()
		db.Take(&user)
	c.JSON(http.StatusOK, user)
}