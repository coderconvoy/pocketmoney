package main

import (
	"fmt"
	"net/http"
	"strings"
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
	Jobs []JPar
}

// SetJob is intended to allow this to be passed around in
func (pd *PageData) SetJob(k string, v interface{}) *PageData {
	for i, j := range pd.Jobs {
		if j.s == k {
			pd.Jobs[i] = JPar{k, v}
			return pd
		}
	}
	pd.Jobs = append(pd.Jobs, JPar{k, v})
	return pd
}

func (pd PageData) Job(k string) interface{} {
	for _, v := range pd.Jobs {
		if v.s == k {
			return v.i
		}
	}
	return nil
}

func NewPageData(mes, fmem string, fam *Family) *PageData {
	return &PageData{
		Mes:  mes,
		Fmem: fmem,
		Fam:  fam,
		Jobs: []JPar{},
	}
}

type LoginData struct {
	W      http.ResponseWriter
	R      *http.Request
	Fam    *Family
	Fmem   string
	LockID uint64
}

type JPar struct {
	s string
	i interface{}
}

func (ld LoginData) Pd(mes string, js ...JPar) *PageData {
	return &PageData{
		Mes:  mes,
		Fam:  ld.Fam,
		Fmem: ld.Fmem,
		Jobs: js,
	}
}

type ACPageData struct {
	Fmem string
	AC   Account
	List []Transaction
	RT   []int
}

type Family struct {
	FamilyName           string
	Members              []User
	Accounts             []*Account
	Transactions         []Transaction
	Requests             []Transaction
	Standing             []StandingOrder
	LastCalc, LastChange time.Time
}

type User struct {
	Username string
	Email    string
	Password dbase2.Password
	Parent   bool
}

type BasicTransaction struct {
	From, Dest ACKey
	Amount     int
	Purpose    string
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
	Start, Stop  time.Time
	Rules        string
	Interval     int
	IntervalType int
}

type ACKey struct {
	Username, Name string
}

func NewACKey(s string) (ACKey, error) {
	sp := strings.Split(s, ":")
	if len(sp) != 2 {
		return ACKey{}, fmt.Errorf("Could not Parse '%s'.", s)
	}
	return ACKey{sp[0], sp[1]}, nil
}

func (a ACKey) String() string { return a.Username + ":" + a.Name }

type Account struct {
	ACKey
	StartDate time.Time
	Current   int
	Available int
}
