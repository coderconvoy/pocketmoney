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
	//TODO, check permission to add account

	uname := r.FormValue("username")
	aname := r.FormValue("accountname")

	fam.Accounts = append(fam.Accounts, Account{
		Username:  uname,
		ID:        len(fam.Accounts), //TODo generate
		Name:      aname,
		StartDate: time.Now(),
		Current:   0,
	})
	err := SaveFamily(fam)

	if err != nil {
		fam.Accounts = fam.Accounts[:len(Fam.Accounts)-1]

	}

}
