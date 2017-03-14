package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/coderconvoy/pocketmoney/history"
	"github.com/pkg/errors"
)

func readPostTransaction(ld LoginData) (history.Transaction, error) {
	r := ld.R
	res := history.Transaction{}
	var err error

	res.From, err = history.NewACKey(r.FormValue("from"))
	if err != nil {
		return res, errors.Wrap(err, "From account not parseable")
	}

	res.Dest, err = history.NewACKey(r.FormValue("to"))
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

	bt, err := readPostTransaction(ld)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd(err.Error()))
		return
	}
	bt.Date = time.Now()

	err = fam.Period.ApplyTransaction(bt)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd(err.Error()))
		return
	}

	ExTemplate(GT, w, "userhome.html", ld.Pd(""))
}

func HandleTransactions(ld LoginData) {
	ExTemplate(GT, ld.W, "transactions.html", ld.Pd(""))
}

func HandleViewAccount(ld LoginData) {
	w, r, fam, fmem := ld.W, ld.R, ld.Fam, ld.Fmem
	rname := r.FormValue("uname")
	rac := r.FormValue("ac")
	//parent or own allowed
	if !fam.IsParent(fmem) && fmem != rname {
		ExTemplate(GT, w, "userhome.html", ld.Pd("Must be your own page"))

		return
	}
	fam.Calculate()

	ExTemplate(GT, w, "viewac.html", ld.Pd("", JPar{"ac", history.ACKey{rname, rac}}))
}

func HandleAddStanding(ld LoginData) {
	w, r, fam := ld.W, ld.R, ld.Fam

	bt, err := readPostTransaction(ld)
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

	nstand := &StandingOrder{
		Transaction:  bt,
		Start:        stime,
		Interval:     delay,
		IntervalType: delayType,
	}
	nstand.Purpose = "$" + nstand.Purpose

	fam.Standing = append(fam.Standing, nstand)

	fam.Calculate()

	ExTemplate(GT, w, "userhome.html", ld.Pd(""))

}
