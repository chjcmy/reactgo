package db

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string
	Content string
	BookName       int
	User User `gorm:"foreignkey:BookName;"`
	BookCode int
	BookSubject BookSubject `gorm:"foreignkey:BookCode;"`
}

type User struct {
	gorm.Model
	UserName string
	Email string
	Github string
	Gitlab string
	Manager int
}

type BookSubject struct {
	gorm.Model
	BookSubjectCode string
	Subject string
}