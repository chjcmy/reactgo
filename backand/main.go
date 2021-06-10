package main

import (
db "backand/db"

)

type Product struct {
	Code  string
	Price uint
}

func main() {
	db := db.Connect()

	db.AutoMigrate(&Product{})
}