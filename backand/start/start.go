package main

import (
	"github.com/backand/ent"
	"github.com/backand/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"log"
)

func main() {
	client, err := ent.Open("mysql", "cshcmi:chltjdgus123!@tcp(192.168.0.8:1994)/blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Panicf("failed creating schema resources: %v", err)
	}
}
