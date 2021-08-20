package db

import (
	"github.com/backand/ent"
)

func Config() *ent.Client {
	client, _ := ent.Open("mysql", "cshcmi:chltjdgus123!@tcp(localhost:1994)/blog?charset=utf8mb4&parseTime=True")
	return client
}
