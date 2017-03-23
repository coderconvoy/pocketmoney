package main

import (
	"time"

	"github.com/coderconvoy/dbase2"
	"github.com/coderconvoy/pocketmoney/history"
)

func HandleAddMember(ld *PageHand) (string, string) {
	r, fam, fmem := ld.R, ld.Fam, ld.Fmem
	if !fam.IsParent(fmem) {
		return "/family", "Not a Parent"
	}

	uname := r.FormValue("username")
	parent := r.FormValue("parent")
	pwd1 := r.FormValue("pwd1")
	pwd2 := r.FormValue("pwd2")

	if uname == "" {
		return "/family", "No Username"
	}
	if pwd1 != pwd2 || pwd1 == "" {
		return "/family", "Passwords not matching"
	}

	pw, err := dbase2.NewPassword(pwd1)
	if err != nil {
		return "/family", "Could not Password: " + err.Error()
	}

	for _, m := range fam.Members {
		if m.Username == uname {
			return "/family", "Username already in use"
		}
	}
	fam.Members = append(fam.Members, User{
		Username: uname,
		Parent:   parent == "on",
		Password: pw,
	})
	fam.Period.Accounts = append(fam.Period.Accounts, &history.Account{
		ACKey:  history.ACKey{uname, "checking"},
		Opened: time.Now(),
	})

	return "/family", ""
}
