package main

import (
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

	amount, err := history.ParseAmount(r.FormValue("amount"))
	if err != nil {
		return res, errors.Wrap(err, "Cannot parse amount")

	}

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

	ExTemplate(GT, w, "viewac.html", ld.Pd("", JPar{"ac", history.ACKey{rname, rac}}))
}
