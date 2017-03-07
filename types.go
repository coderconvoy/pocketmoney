package main

import (
	"time"

	"github.com/coderconvoy/dbase2"
)

const (
	T_PAID = iota
	T_REQUESTED
	T_REJECTED
)
const (
	D_NDAYS = iota
	D_OFMONTH
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
	Requests     []Transaction
	Standing     []StandingOrder
	LastStanding time.Time
}

type User struct {
	Username string
	Email    string
	Password dbase2.Password
	Parent   bool
}

type BasicTransaction struct {
	FromUser, DestUser string
	FromAC, DestAC     string
	Amount             int
	Purpose            string
}

type Transaction struct {
	BasicTransaction
	Status int
	Date   time.Time
}

type StandingOrder struct {
	BasicTransaction
	Start     time.Time
	Delay     int
	DelayType int
}

type Account struct {
	Username  string
	Name      string
	StartDate time.Time
	Current   int
	Available int
}
