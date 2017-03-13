package main

import (
	"errors"
	"time"

	"github.com/coderconvoy/pocketmoney/history"
)

func (f *Family) Calculate() {
	if f.LastCalc.After(f.LastChange) && time.Now().After(f.LastCalc.Add(time.Hour*5)) {
		return
	}
	//	f.CalculateStanding()

}

func (f *Family) calculateTransaction(t history.Transaction) error {

	fac, fok := f.Account(t.From)
	dac, dok := f.Account(t.Dest)

	if !(fok && dok) {
		return errors.New("No Account for Transaction")
	}
	fac.Current -= t.Amount
	dac.Current += t.Amount
	return nil
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

func (f *Family) AccumulateTransactions(ak history.ACKey) []history.Accumulation {
	return f.Period.Accumulate(ak)

}
