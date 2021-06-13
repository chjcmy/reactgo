package api

import (
	db3 "backand/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Unswerving struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func Userinfo(c *gin.Context) {

	var user db3.User
		req := &Unswerving{}
		err := c.Bind(req)
		db := db3.Connect()
		 db.Find(&user, "user_name = ?", req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}