package main

import (
	"./ent"
	"log"
)

func main() {
	client, err := ent.Open("mysql", "cshcmi:chltjdgus123!@tcp(58.229.145.190:1994)/book")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
}
