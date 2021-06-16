package api

import (
	db3 "backand/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type BookResult struct {
	Id         uint
	CreateDate time.Time
	UpdateDate time.Time
	Title      string
	Content    string
	Subject    string
}

type BookTitleResult struct {
	Id         uint
	CreateDate time.Time
	UpdateDate time.Time
	Title      string
	Subject    string
}

type BookId struct {
	ID int `uri:"id" binding:"required"`
}

func ShowSubject(c *gin.Context) {
	var subject []db3.BookSubject
	db := db3.Connect()
	db.Find(&subject)
	c.JSON(http.StatusOK, subject)
}

func ShowBook(c *gin.Context) {
	db := db3.Connect()
	var booked BookId
	if err := c.ShouldBindUri(&booked); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	var result BookResult
	db.Table("books").
		Select("books.id , books.created_at as create_date, books.updated_at as update_date, books.title , books.content , book_subjects.subject").
		Joins("left join book_subjects on books.book_code = book_subjects.id").
		Where("books.id = ?", booked.ID).
		Scan(&result)
	c.JSON(http.StatusOK, &result)
}

func ShowTitles(c *gin.Context) {
	var result []BookTitleResult
	db := db3.Connect()
	db.Table("books").
		Select("books.id, books.created_at as create_date, books.updated_at as update_date, books.title , book_subjects.subject").
		Joins("left join book_subjects on books.book_code = book_subjects.id").
		Scan(&result)
	c.JSON(http.StatusOK, &result)
}
