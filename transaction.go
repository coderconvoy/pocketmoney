package main

import (
	"time"

	"github.com/coderconvoy/pocketmoney/history"
	"github.com/pkg/errors"
)

func readPostTransaction(ld *PageHand) (history.Transaction, error) {

	r := ld.R
	res := history.Transaction{}
	var err error

	res.From.Username = r.FormValue("fromuser")
	if res.From.Username == "" {

		res.From, err = history.NewACKey(r.FormValue("from"))
		if err != nil {
			return res, errors.Wrap(err, "From account not parseable")
		}

	}

	res.Dest, err = history.NewACKey(r.FormValue("to"))
	if err != nil {
		return res, errors.Wrap(err, "Dest account not parseable")
	}
	if res.From == res.Dest {
		return res, errors.New("From and Destination are the same account")
	}
	//check dest and from have real accounts
	fnd, dnd := false, false
	for _, a := range ld.Fam.Period.Accounts {
		if a.ACKey == res.Dest {
			dnd = true
		}
		if a.ACKey == res.From {
			fnd = true
		}
		if a.Username == res.From.Username && res.From.Name == "" && !fnd {
			fnd = true
			res.From.Name = a.Name //Must have account name, for safety
		}
	}
	if !fnd || !dnd {
		return res, errors.New("Accounts must be real, no hacky hacky")
	}

	res.Amount, err = history.ParseAmount(r.FormValue("amount"))
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

func HandlePay(ld *PageHand) (string, string) {
	fam := ld.Fam

	bt, err := readPostTransaction(ld)
	if err != nil {
		return "/personal", err.Error()
	}
	bt.Date = time.Now()

	err = fam.Period.ApplyTransaction(bt)
	if err != nil {
		return "/personal", err.Error()
	}

	return "/personal", ""
}

func HandleViewAccount(ld *PageHand) string {
	r, fam, fmem := ld.R, ld.Fam, ld.Fmem
	rname := r.FormValue("uname")
	rac := r.FormValue("ac")
	//parent or own allowed
	if !fam.IsParent(fmem) && fmem != rname {
		ld.Mes = "Must be your own page"
		return "userhome.html"
	}

	ld.SetJob("ac", history.ACKey{rname, rac})
	return "viewac.html"
}
