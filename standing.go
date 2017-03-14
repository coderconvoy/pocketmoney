package main

import (
	"time"

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
