package main

import (
	"net/http"
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
	Jobs []string
	JobN int
}

func NewPageData(mes, fmem string, fam *Family) PageData {
	return PageData{
		Mes:  mes,
		Fmem: fmem,
		Fam:  fam,
		Jobs: []string{},
		JobN: 0,
	}
}

type LoginData struct {
	W      http.ResponseWriter
	R      *http.Request
	Fam    *Family
	Fmem   string
	LockID uint64
}

func (ld LoginData) Pd(s ...string) PageData {
	if len(s) == 0 {
		s = []string{""}
	}
	return PageData{
		Mes:  s[0],
		Fam:  ld.Fam,
		Fmem: ld.Fmem,
		Jobs: s[1:],
		JobN: 0,
	}
}

type ACPageData struct {
	Fmem string
	AC   Account
	List []Transaction
	RT   []int
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

type Transortable []Transaction

func (t Transortable) Len() int           { return len(t) }
func (t Transortable) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Transortable) Less(i, j int) bool { return t[j].Date.After(t[i].Date) }

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
