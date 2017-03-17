package main

import (
	"strconv"
	"time"

	"github.com/coderconvoy/dbase2"
	"github.com/coderconvoy/pocketmoney/history"
)

func (so *StandingOrder) Next() (history.Transaction, bool) {
	nd := time.Time{}
	if so.Date == nd {
		nd = so.Start
	} else {
		nd = NextDate(so.Date, so.Interval, so.IntervalType)
	}
	if so.Start.After(nd) || nd.After(time.Now()) ||
		(so.Stop.After(time.Time{}) && nd.After(so.Stop)) {
		return so.Transaction, false
	}
	so.Date = nd
	return so.Transaction, true
}

func NextDate(d time.Time, step int, steptype int) time.Time {
	if step <= 0 {
		step = 1
	}
	if steptype == D_NDAYS {
		return d.AddDate(0, 0, step)
	}
	return d.AddDate(0, step, 0)
}

func HandleAddStanding(ld LoginData) {
	w, r, fam := ld.W, ld.R, ld.Fam

	bt, err := readPostTransaction(ld)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd(err.Error()))
		return
	}

	start := r.FormValue("start")
	dbase2.QLog("start:" + start)
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

	//lazy, fix if new types are added
	delayType := D_NDAYS
	if r.FormValue("delaytype") == "days" {
		delayType = D_OFMONTH
	}

	nstand := &StandingOrder{
		Transaction:  bt,
		Start:        stime,
		Interval:     delay,
		IntervalType: delayType,
		ID:           fam.NewStandingID(),
	}
	nstand.Purpose = "$" + nstand.Purpose

	fam.Standing = append(fam.Standing, nstand)

	fam.Calculate()

	ExTemplate(GT, w, "userhome.html", ld.Pd(""))
}

func HandleCancelStanding(ld LoginData) {
	rmid64, err := strconv.ParseInt(ld.R.FormValue("id"), 10, 32)
	if err != nil {
		ExTemplate(GT, ld.W, "userhome.html", ld.Pd("No id Given"))
		return
	}
	rmid := int32(rmid64)

	fnd := -1
	for i, o := range ld.Fam.Standing {
		if o.ID == rmid {
			fnd = i
		}
	}
	if fnd < 0 {
		ExTemplate(GT, ld.W, "userhome.html", ld.Pd("No order matches ID given"))
		return
	}

	ld.Fam.Standing = append(ld.Fam.Standing[:fnd], ld.Fam.Standing[fnd+1:]...)

	ExTemplate(GT, ld.W, "userhome.html", ld.Pd(""))

}
