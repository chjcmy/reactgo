package api

import (
	db3 "backand/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type BookResult struct {
	Id         uint
	CreateDate time.Time
	UpdateDate time.Time
	Title      string
	Content    string
	Name       string
	Subject    string
}

type BookTitleResult struct {
	Id         uint
	CreateDate time.Time
	UpdateDate time.Time
	Title      string
	Name       string
	Subject    string
}

type BookId struct {
	ID int `uri:"id" binding:"required"`
}

type InsertValue struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	SubjectId int    `json:"subject"`
	Name      string `header:"Name"`
}

type UpdateValue struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
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
		Select("books.id , books.created_at as create_date, books.updated_at as update_date, books.title , books.content , book_subjects.subject, users.user_name as name").
		Joins("left join book_subjects on books.book_code = book_subjects.id").
		Joins("left join users on books.book_name = users.id").
		Where("books.id = ?", booked.ID).
		Scan(&result)
	c.JSON(http.StatusOK, &result)
}

func ShowTitles(c *gin.Context) {
	var result []BookTitleResult
	db := db3.Connect()
	db.Table("books").
		Select("books.id, books.created_at as create_date, books.updated_at as update_date, books.title , book_subjects.subject, users.user_name as name").
		Joins("left join book_subjects on books.book_code = book_subjects.id").
		Joins("left join users on books.book_name = users.id").
		Scan(&result)
	c.JSON(http.StatusOK, &result)
}

func SubjectTitles(c *gin.Context) {
	var subjected BookId
	if err := c.ShouldBindUri(&subjected); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	var result []BookTitleResult
	db := db3.Connect()
	db.Table("books").
		Select("books.id, books.created_at as create_date, books.updated_at as update_date, books.title , book_subjects.subject, users.user_name as name").
		Joins("left join book_subjects on books.book_code = book_subjects.id").
		Joins("left join users on books.book_name = users.id").
		Where("books.book_code = ?", subjected.ID).
		Scan(&result)
	c.JSON(http.StatusOK, &result)
}

func CreateBook(c *gin.Context) {
	var result InsertValue
	var named BookId
	err := c.ShouldBind(&result)
	if err != nil {
		return
	}
	err2 := c.ShouldBindHeader(&result)
	if err2 != nil {
		return
	}
	db := db3.Connect()
	db.Table("users").
		Select("users.id").
		Where("users.email = ?", result.Name).
		Scan(&named)
	fmt.Println(&named.ID)
	value := db3.Book{
		Model:       gorm.Model{},
		Title:       result.Title,
		Content:     result.Content,
		BookName:    named.ID,
		User:        db3.User{},
		BookCode:    named.ID,
		BookSubject: db3.BookSubject{},
	}
	db.Create(&value)

	c.JSON(http.StatusOK, &result)
}

func UpdateBook(c *gin.Context) {
	var result UpdateValue
	err := c.ShouldBind(&result)
	if err != nil {
		return
	}
	db := db3.Connect()
	value := db3.Book{
		Model:   gorm.Model{},
		Title:   result.Title,
		Content: result.Content,
	}
	db.Model(&db3.Book{}).
		Where("id = ?", result.Id).
		Updates(&value)

	c.JSON(http.StatusOK, &result)
}

func DeleteBook(c *gin.Context) {
	var booked BookId
	if err := c.ShouldBindUri(&booked); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	db := db3.Connect()
	db.Unscoped().Where("id = ?", booked.ID).Delete(&db3.Book{})

	c.JSON(http.StatusOK, "success")
}
