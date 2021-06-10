package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "cshcmi:chltjdgus123!@tcp(58.229.145.190:1994)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		println("연결 성공")
	}
	if err != nil {
		panic(err)
	}

	return db
}
