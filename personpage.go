package main

import (
	"github.com/coderconvoy/pocketmoney/history"
)

func HandleAddAccount(ld *PageHand) (string, string) {
	r, fam := ld.R, ld.Fam
	//TODO, check permission to add account

	aname := r.FormValue("accountname")

	ac := history.CreateAccount(ld.Fmem, aname)
	ac.Col1 = r.FormValue("Col1")
	ac.Col2 = r.FormValue("Col2")
	fam.Period.Accounts = append(fam.Period.Accounts, ac)

	return "/personal", ""
}
