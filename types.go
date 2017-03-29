package main

import (
	"net/http"
	"time"

	"github.com/coderconvoy/dbase"
	"github.com/coderconvoy/pocketmoney/history"
)

const (
	D_NDAYS = iota
	D_OFMONTH
)

type LoginPart struct {
	Fam  string
	User string
}

type LoginStore struct {
	Familyname string
	Fmem       string
	Jobs       []JPar
	Mes        interface{}
}

type PageData struct {
	Fam *Family
	LoginStore
}

type PageHand struct {
	*PageData
	W http.ResponseWriter
	R *http.Request
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

func JPars(js ...string) []JPar {
	if len(js) == 1 {
		return []JPar{{"mes", js[0]}}
	}
	res := []JPar{}
	for i := 1; i < len(js); i += 2 {
		res = append(res, JPar{js[i-1], js[i]})
	}
	return res
}

type ACPageData struct {
	Fmem string
	AC   history.Account
	List []history.Transaction
	RT   []int
}

type Family struct {
	history.Period       `json:"Period"`
	FamilyName           string
	Members              []User
	Requests             []*PaymentRequest
	Standing             []*StandingOrder
	LastCalc, LastChange time.Time
}

type User struct {
	Username string
	Email    string
	Password dbase.Password
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
