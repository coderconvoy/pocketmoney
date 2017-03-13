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
