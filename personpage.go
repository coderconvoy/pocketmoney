package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandlePersonal(w http.ResponseWriter, r *http.Request) {
	fam, uname, err := LoggedInFamily(w, r)

	if err != nil {
		GoIndex(w, r, err.Error())
		return
	}

	pdat := PageData{
		Mes:  "",
		Fmem: uname,
		Fam:  fam,
	}
	fam.Calculate()

	ExTemplate(GT, w, "userhome.html", pdat)
}

func HandleAddAccount(w http.ResponseWriter, r *http.Request) {
	fam, fmem, err := LoggedInFamily(w, r)
	if err != nil {
		GoIndex(w, r, err.Error())
		return
	}
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
	err = SaveFamily(fam)

	mes := ""
	if err != nil {
		fam.Accounts = fam.Accounts[:len(fam.Accounts)-1]
		mes = "Could not Save Family: " + err.Error()
	}

	ExTemplate(GT, w, "userhome.html", PageData{mes, fmem, fam})
}

func HandlePay(w http.ResponseWriter, r *http.Request) {
	fam, fmem, err := LoggedInFamily(w, r)
	if err != nil {
		GoIndex(w, r, err.Error())
		return
	}
	fUser := r.FormValue("username")
	fAcc := r.FormValue("from")
	toData := r.FormValue("to")
	toSpl := strings.Split(toData, ":")
	if len(toSpl) != 2 {
		ExTemplate(GT, w, "userhome.html", PageData{"No ':' error", fmem, fam})
		return
	}

	amount := r.FormValue("amount")
	am, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", PageData{"Amount must be a number", fmem, fam})
		return
	}

	purpose := r.FormValue("purpose")

	fam.Transactions = append(fam.Transactions, Transaction{
		BasicTransaction: BasicTransaction{
			FromUser: fUser,
			FromAC:   fAcc,
			DestUser: toSpl[0],
			DestAC:   toSpl[1],
			Amount:   int(am * 100),
			Purpose:  purpose,
		},
		Status: T_PAID,
		Date:   time.Now(),
	})
	fam.Calculate()
	err = SaveFamily(fam)
	mes := ""
	if err != nil {
		fam.Accounts = fam.Accounts[:len(fam.Accounts)-1]
		mes = "Could not Save Family: " + err.Error()
	}

	ExTemplate(GT, w, "userhome.html", PageData{mes, fmem, fam})
}
func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	fam, fmem, err := LoggedInFamily(w, r)
	if err != nil {
		GoIndex(w, r, err.Error())
		return
	}
	ExTemplate(GT, w, "transactions.html", PageData{"", fmem, fam})
}

func HandleViewAccount(w http.ResponseWriter, r *http.Request) {
	fam, fmem, err := LoggedInFamily(w, r)
	if err != nil {
		GoIndex(w, r, err.Error())
		return
	}

	rname := r.FormValue("uname")
	rac := r.FormValue("ac")
	//parent or own allowed
	if !IsParent(fmem, fam) || fmem != rname {
		ExTemplate(GT, w, "userpage.html", PageData{"Must be your own page", fmem, fam})

		return
	}
	fam.Calculate()
	ac, tList, rt, err := fam.ACTransactions(rname, rac)
	if err != nil {
		ExTemplate(GT, w, "userpage.html", PageData{err.Error(), fmem, fam})

		return
	}
	ExTemplate(GT, w, "viewac.html", ACPageData{fmem, ac, tList, rt})
}
