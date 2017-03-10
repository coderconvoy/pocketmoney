package main

import (
	"time"
)

func HandlePersonal(ld LoginData) {
	w, fam, uname := ld.W, ld.Fam, ld.Fmem

	pdat := PageData{
		Mes:  "",
		Fmem: uname,
		Fam:  fam,
	}
	fam.Calculate()

	ExTemplate(GT, w, "userhome.html", pdat)
}

func HandleAddAccount(ld LoginData) {
	w, r, fam := ld.W, ld.R, ld.Fam
	//TODO, check permission to add account

	uname := r.FormValue("username")
	aname := r.FormValue("accountname")

	fam.Accounts = append(fam.Accounts, &Account{
		Username:  uname,
		Name:      aname,
		StartDate: time.Now(),
		Current:   0,
	})

	fam.Calculate()
	err := SaveFamily(fam)

	mes := ""
	if err != nil {
		fam.Accounts = fam.Accounts[:len(fam.Accounts)-1]
		mes = "Could not Save Family: " + err.Error()
	}

	ExTemplate(GT, w, "userhome.html", ld.Pd(mes))
}
