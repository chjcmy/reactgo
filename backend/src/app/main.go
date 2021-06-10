package main

import (
	db "./src/db"

)

type Product struct {
	Code  string
	Price uint
}

func main() {
	db := db.Connect()

	db.AutoMigrate(&Product{})
}