package main

import (
	"strconv"
	"time"

	"github.com/coderconvoy/dbase"
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

func HandleAddStanding(ld *PageHand) (string, string) {
	r, fam := ld.R, ld.Fam

	bt, err := readPostTransaction(ld)
	if err != nil {
		return "/personal", err.Error()
	}

	start := r.FormValue("start")
	dbase.QLog("start:" + start)
	stime, err := time.Parse("2006-01-02", start)
	if err != nil {
		return "/personal", "could not parse date:" + err.Error()
	}

	if time.Now().AddDate(0, 0, -1).After(stime) {
		return "/personal", "Must be a future date"
	}

	delay, err := strconv.Atoi(r.FormValue("delay"))
	if err != nil {
		return "/personal", "could not parse delay interval"
	}
	if delay < 1 {
		return "/personal", "interval must be positive"
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

	return "/personal", ""
}

func HandleCancelStanding(ld *PageHand) (string, string) {
	rmid64, err := strconv.ParseInt(ld.R.FormValue("id"), 10, 32)
	if err != nil {
		return "/personal", "No Id Given"
	}
	rmid := int32(rmid64)

	fnd := -1
	for i, o := range ld.Fam.Standing {
		if o.ID == rmid {
			fnd = i
		}
	}
	if fnd < 0 {
		return "/personal", "No Order matches ID given"
	}

	ld.Fam.Standing = append(ld.Fam.Standing[:fnd], ld.Fam.Standing[fnd+1:]...)

	return "/personal", ""

}
