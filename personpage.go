package main

import (
	"github.com/coderconvoy/pocketmoney/history"
)

func HandlePersonal(ld LoginData) {
	w, fam := ld.W, ld.Fam

	fam.Calculate()

	ExTemplate(GT, w, "userhome.html", ld.Pd(""))
}

func HandleAddAccount(ld LoginData) {
	w, r, fam := ld.W, ld.R, ld.Fam
	//TODO, check permission to add account

	uname := r.FormValue("username")
	aname := r.FormValue("accountname")

	fam.Period.Accounts = append(fam.Period.Accounts, history.CreateAccount(uname, aname))

	ExTemplate(GT, w, "userhome.html", ld.Pd(""))
}
