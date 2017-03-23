package main

import (
	"github.com/coderconvoy/pocketmoney/history"
)

func HandleAddAccount(ld *PageHand) (string, string) {
	r, fam := ld.R, ld.Fam
	//TODO, check permission to add account

	aname := r.FormValue("accountname")

	fam.Period.Accounts = append(fam.Period.Accounts, history.CreateAccount(ld.Fmem, aname))

	return "/personal", ""
}
