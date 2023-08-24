package main

import(
	"log"
)

func seedAccount(store Storage, fname, lname, pw String) *Account{
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	return acc
}

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8001", store)
	server.Run()
}
