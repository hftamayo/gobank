package main

import(
	"log"
	"fmt"
	"flag"
)

func seedAccount(store Storage, fname, lname, pw string) *Account{
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)
	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Michael", "Jordan", "MJ123")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed{
		fmt.Println("seeding the database...")
		seedAccounts(store)
	}


	server := NewAPIServer(":8001", store)
	server.Run()
}
