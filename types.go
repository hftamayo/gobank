package main

import "math/rand"

type Account struct {
	ID			int `json:"id"`
	FirstName	string `json:"first_name"`
	LastName	string `json:"last_name"`
	Number		int64 `json:"Bank_Account"`
	Balance 	int64 `json:"Current_Balance"`
}

func NewAccount(firstName, LastName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: LastName,
		Number: int64(rand.Intn(1000000)),
	}
}