package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func readPostBasicTransaction(ld LoginData) (BasicTransaction, error) {
	r := ld.R
	fData := r.FormValue("from")
	toData := r.FormValue("to")
	res := BasicTransaction{}

	if fData == toData {
		return res, errors.New("From and Destination are the same")
	}

	fSpl := strings.Split(fData, ":")
	if len(fSpl) != 2 {
		return res, errors.New("From account not parseable")
	}
	res.FromUser = fSpl[0]
	res.FromAC = fSpl[1]

	toSpl := strings.Split(toData, ":")
	if len(toSpl) != 2 {
		return res, errors.New("To account not parseable")
	}
	res.DestUser = toSpl[0]
	res.DestAC = toSpl[1]

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
	w, fam, fmem := ld.W, ld.Fam, ld.Fmem

	bt, err := readPostBasicTransaction(ld)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", PageData{err.Error(), fmem, fam})
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

	ExTemplate(GT, w, "userhome.html", PageData{mes, fmem, fam})
}

func HandleTransactions(ld LoginData) {
	w, fam, fmem := ld.W, ld.Fam, ld.Fmem
	ExTemplate(GT, w, "transactions.html", PageData{"", fmem, fam})
}

func HandleViewAccount(ld LoginData) {
	w, r, fam, fmem := ld.W, ld.R, ld.Fam, ld.Fmem
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

func HandleAddStanding(ld LoginData) {
	w, r, fam, fmem := ld.W, ld.R, ld.Fam, ld.Fmem

	bt, err := readPostBasicTransaction(ld)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", PageData{err.Error(), fmem, fam})
		return
	}

	start := r.FormValue("start")
	fmt.Println("start:", start)
	stime, err := time.Parse("2006-01-02", start)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", PageData{"could not parse date" + err.Error(), fmem, fam})
		return
	}

	if time.Now().AddDate(0, 0, -1).After(stime) {
		ExTemplate(GT, w, "userhome.html", PageData{"Must be a future startdate", fmem, fam})
	}

	delay, err := strconv.Atoi(r.FormValue("delay"))
	if err != nil {
		ExTemplate(GT, w, "userhome.html", PageData{"could not parse delay interval" + err.Error(), fmem, fam})
		return
	}
	if delay < 1 {
		ExTemplate(GT, w, "userhome.html", PageData{"Must increment positive", fmem, fam})
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

	fam.Standing = append(fam.Standing, nstand)
	fam.Calculate()
	err = SaveFamily(fam)
	mes := ""
	if err != nil {
		fam.Accounts = fam.Accounts[:len(fam.Accounts)-1]
		mes = "Could not Save Family: " + err.Error()
	}

	ExTemplate(GT, w, "userhome.html", PageData{mes, fmem, fam})

}
