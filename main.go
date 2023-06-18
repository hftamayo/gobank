package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":8001")
	server.Run()
}