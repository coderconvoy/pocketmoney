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
	FromUser, DestUser string
	FromID, DestID     string
	Amount             int
	Authorized         bool
	Purpose            string
}

type Account struct {
	Username  string
	ID        uint64
	Name      string
	StartDate time.Time
	Current   int
}
