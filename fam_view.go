package main

import (
	"errors"

	"github.com/coderconvoy/pocketmoney/history"
)

func (fam *Family) GetUser(uname string) (*User, error) {
	for i, m := range fam.Members {
		if m.Username == uname {
			return &fam.Members[i], nil
		}
	}
	return nil, errors.New("No Member of that name")
}

func (fam *Family) IsParent(uname string) bool {
	m, err := fam.GetUser(uname)
	if err != nil {
		return false
	}
	return m.Parent
}

func (fam *Family) WriteAccess(a *history.Account, uname string) bool {
	if a.Username == uname {
		return true
	}
	return a.Username == "WORLD" && fam.IsParent(uname)
}

func (fam *Family) ListWriteAccess(uname string) []*history.Account {
	res := []*history.Account{}
	for _, v := range fam.Period.Accounts {
		if fam.WriteAccess(v, uname) {
			res = append(res, v)
		}
	}
	return res
}

func (fam *Family) FilterStandingByUser(uname string) []*StandingOrder {
	res := []*StandingOrder{}
	isPar := fam.IsParent(uname)
	for _, v := range fam.Standing {
		if v.From.Username == uname ||
			(v.From.Username == "WORLD" && isPar) {
			res = append(res, v)
		} else if v.Dest.Username == uname ||
			(v.Dest.Username == "WORLD" && isPar) {
			res = append(res, v)
		}
	}
	return res
}

func (fam *Family) CanEditStanding(so *StandingOrder, uname string) bool {
	return so.From.Username == uname ||
		(so.From.Username == "WORLD" && fam.IsParent(uname))

}

func (fam *Family) GetMustRequests(uname string) []*PaymentRequest {
	res := []*PaymentRequest{}
	for _, pr := range fam.Requests {
		if (pr.From.Username == uname && pr.Returns%2 == 0) ||
			(pr.Requester == uname && pr.Returns%2 != 0) {
			res = append(res, pr)
		}
	}
	return res

}
