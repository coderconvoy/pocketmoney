package main

import (
	"net/http"
	"time"

	"github.com/coderconvoy/dbase2"
	"github.com/coderconvoy/pocketmoney/history"
)

const (
	D_NDAYS = iota
	D_OFMONTH
)

type LoginStore struct {
	Familyname string
	Fmem       string
	Jobs       []JPar
}

type PageData struct {
	Fam *Family
	LoginStore
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

func JPars(js ...string) {
	if len(js) == 1 {
		return []JPar{{"mes", js[0]}}
	}
	res := []JPar{}
	for i := 1; i < len(js); i += 2 {
		res = append(res, JPar{js[i-1], js[i]})
	}
	return res
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
	AC   history.Account
	List []history.Transaction
	RT   []int
}

type Family struct {
	FamilyName           string
	Members              []User
	Period               history.Period
	Requests             []*PaymentRequest
	Standing             []*StandingOrder
	LastCalc, LastChange time.Time
}

type User struct {
	Username string
	Email    string
	Password dbase2.Password
	Parent   bool
}

type StandingOrder struct {
	history.Transaction
	Stop, Start  time.Time
	Rules        string
	Interval     int
	IntervalType int
	ID           int32
}

type PaymentRequest struct {
	history.Transaction
	Requester string
	ID        int32
	Returns   int
}
