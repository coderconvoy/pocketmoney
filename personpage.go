package main

import (
	"net/http"
	"time"
)

func HandlePersonal(w http.ResponseWriter, r *http.Request) {
	fam, uname, err := LoggedInFamily(w, r)

	if err != nil {
		ExTemplate(GT, w, "index.html", err.Error())
		return
	}

	pdat := PageData{
		Mes:  "",
		Fmem: uname,
		Fam:  fam,
	}

	ExTemplate(GT, w, "userhome.html", pdat)
}

func HandleAddAccount(w http.ResponseWriter, r *http.Request) {
	fam, fmem, err := LoggedInFamily(w, r)
	if err != nil {
		ExTemplate(GT, w, "index.html", err.Error())
		return
	}
	//TODO, check permission to add account

	uname := r.FormValue("username")
	aname := r.FormValue("accountname")

	fam.Accounts = append(fam.Accounts, Account{
		Username:  uname,
		Name:      aname,
		StartDate: time.Now(),
		Current:   0,
	})
	err = SaveFamily(fam)

	mes := ""
	if err != nil {
		fam.Accounts = fam.Accounts[:len(fam.Accounts)-1]
		mes = "Could not Save Family: " + err.Error()
	}

	ExTemplate(GT, w, "userhome.html", PageData{mes, fmem, fam})
}
