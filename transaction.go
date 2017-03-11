package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func readPostBasicTransaction(ld LoginData) (BasicTransaction, error) {
	r := ld.R
	res := BasicTransaction{}
	var err error

	res.From, err = NewACKey(r.FormValue("from"))
	if err != nil {
		return res, errors.Wrap(err, "From account not parseable")
	}

	res.Dest, err = NewACKey(r.FormValue("to"))
	if err != nil {
		return res, errors.Wrap(err, "Dest account not parseable")
	}

	if res.From == res.Dest {
		return res, errors.New("From and Destination are the same account")
	}

	amount := r.FormValue("amount")
	am, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return res, errors.New("Could not parse amount")
	}
	if am < 0 {
		return res, errors.New("Amount not Positive")
	}
	res.Amount = int(am * 100)

	purpose := r.FormValue("purpose")
	if purpose == "" {
		return res, errors.New("Must have a purpose")
	}
	res.Purpose = purpose
	return res, nil

}

func HandlePay(ld LoginData) {
	w, fam := ld.W, ld.Fam

	bt, err := readPostBasicTransaction(ld)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd(err.Error()))
		return
	}

	fam.Transactions = append(fam.Transactions, Transaction{
		BasicTransaction: bt,
		Status:           T_PAID,
		Date:             time.Now(),
	})
	fam.Calculate()
	err = SaveFamily(fam)
	mes := ""
	if err != nil {
		fam.Accounts = fam.Accounts[:len(fam.Accounts)-1]
		mes = "Could not Save Family: " + err.Error()
	}

	ExTemplate(GT, w, "userhome.html", ld.Pd(mes))
}

func HandleTransactions(ld LoginData) {
	ExTemplate(GT, ld.W, "transactions.html", ld.Pd(""))
}

func HandleViewAccount(ld LoginData) {
	w, r, fam, fmem := ld.W, ld.R, ld.Fam, ld.Fmem
	rname := r.FormValue("uname")
	rac := r.FormValue("ac")
	//parent or own allowed
	if !IsParent(fmem, fam) || fmem != rname {
		ExTemplate(GT, w, "userpage.html", ld.Pd("Must be your own page"))

		return
	}
	fam.Calculate()

	ExTemplate(GT, w, "viewac.html", ld.Pd("", JPar{"ac", ACKey{rname, rac}}))
}

func HandleAddStanding(ld LoginData) {
	w, r, fam := ld.W, ld.R, ld.Fam

	bt, err := readPostBasicTransaction(ld)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd(err.Error()))
		return
	}

	start := r.FormValue("start")
	fmt.Println("start:", start)
	stime, err := time.Parse("2006-01-02", start)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd("could not parse date"+err.Error()))
		return
	}

	if time.Now().AddDate(0, 0, -1).After(stime) {
		ExTemplate(GT, w, "userhome.html", ld.Pd("Must be a future startdate"))
	}

	delay, err := strconv.Atoi(r.FormValue("delay"))
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd("could not parse delay interval"+err.Error()))
		return
	}
	if delay < 1 {
		ExTemplate(GT, w, "userhome.html", ld.Pd("Must increment positive"))
		return
	}

	//lazy fix if new types are added
	delayType := D_NDAYS
	if r.FormValue("delaytype") == "days" {
		delayType = D_OFMONTH
	}

	nstand := StandingOrder{
		BasicTransaction: bt,
		Start:            stime,
		Delay:            delay,
		DelayType:        delayType,
	}
	nstand.Purpose = "$" + nstand.Purpose

	fam.Standing = append(fam.Standing, nstand)
	fam.Calculate()
	err = SaveFamily(fam)
	mes := ""
	if err != nil {
		fam.Accounts = fam.Accounts[:len(fam.Accounts)-1]
		mes = "Could not Save Family: " + err.Error()
	}

	ExTemplate(GT, w, "userhome.html", ld.Pd(mes))

}
