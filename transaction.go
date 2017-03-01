package main

type Transaction struct {
	ID         uint64
	From       uint64
	To         uint64
	Amount     int
	Authorized bool
}

type Account struct {
	ID      uint64
	Name    string
	Monthly []int
}

type User struct {
	UName    string
	FName    string
	Accounts []uint64
}
