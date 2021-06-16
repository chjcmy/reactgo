package db

import (
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserWithJSON struct {
	gorm.Model
	Name       string
	Attributes datatypes.JSON
}

func Connect() *gorm.DB {
	dsn := "cshcmi:chltjdgus123!@tcp(choi1994.iptime.org:1994)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		println("연결 성공")
	}
	if err != nil {
		panic(err)
	}

	return db
}
