package main

import (
	"time"

	"github.com/coderconvoy/dbase2"
)

type PageData struct {
	Mes  string
	Fmem string
	Fam  *Family
}

type Family struct {
	FamilyName   string
	Members      []User
	Accounts     []*Account
	Transactions []Transaction
}

type User struct {
	Username string
	Email    string
	Password dbase2.Password
	Parent   bool
}

type Transaction struct {
	FromUser, DestUser string
	FromAC, DestAC     string
	Amount             int
	Authorised         bool
	Purpose            string
	Date               time.Time
}

type Account struct {
	Username  string
	Name      string
	StartDate time.Time
	Current   int
}
