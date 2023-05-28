package main

import "math/rand"

type Account struct {
	ID			int `json: "id"`
	FirstName	string `json: "firstName"`
	LastName	string `json: "lastName"`
	Number		int64 `json: "Bank Account"`
	Balance 	int64 `json: "Current Balance"`
}

func NewAccount(firstName, LastName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: LastName,
		Number: int64(rand.Intn(1000000)),
	}
}