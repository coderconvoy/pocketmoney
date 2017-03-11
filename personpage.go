package main

import (
	"time"
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

	fam.Accounts = append(fam.Accounts, &Account{
		ACKey:     ACKey{uname, aname},
		StartDate: time.Now(),
		Current:   0,
	})

	fam.Calculate()
	err := fam.Save()

	mes := ""
	if err != nil {
		fam.Accounts = fam.Accounts[:len(fam.Accounts)-1]
		mes = "Could not Save Family: " + err.Error()
	}

	ExTemplate(GT, w, "userhome.html", ld.Pd(mes))
}
