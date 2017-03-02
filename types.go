package main

import "github.com/coderconvoy/dbase2"

type Family struct {
	FamilyName string
	Members    []User
	Accounts   []Account
}

type User struct {
	Username string
	Email    string
	Password dbase2.Password
	Parent   bool
}

type Transaction struct {
	ID                 uint64
	FromUser, DestUser string
	FromID, DestID     int
	Amount             int
	Authorized         bool
	Purpose            string
}

type Account struct {
	Username string
	ID       uint64
	Name     string
	Monthly  []int
}
