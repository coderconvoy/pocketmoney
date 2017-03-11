package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func (f *Family) Calculate() {
	if f.LastCalc.After(f.LastChange) && time.Now().After(f.LastCalc.Add(time.Hour*5)) {
		return
	}
	f.CalculateStanding()
	f.CalculateTransactions()

	f.LastCalc = time.Now()
}

func (f *Family) CalculateTransactions() {
	fmt.Println("Calculating")
	for _, a := range f.Accounts {
		a.Current = 0
	}

	for _, t := range f.Transactions {
		f.calculateTransaction(t)
		//TODO, work out what to do with error transactions
	}
}

func (f *Family) calculateTransaction(t Transaction) error {

	fac, fok := f.Account(t.From)
	dac, dok := f.Account(t.Dest)

	if !(fok && dok) {
		return errors.New("No Account for Transaction")
	}
	fac.Current -= t.Amount
	dac.Current += t.Amount
	return nil
}

func (f *Family) CalculateStanding() {
	f.calculateStanding(time.Now())
}

func (f *Family) calculateStanding(now time.Time) {
	ntList := []Transaction{}
	for _, s := range f.Standing {

		//Some standing orders will be set in the future.
		if s.Start.After(now) {
			continue
		}

		var lastTrans Transaction

		for i := len(f.Transactions) - 1; i >= 0; i-- {
			if f.Transactions[i].Purpose == s.Purpose {
				lastTrans = f.Transactions[i]
				break
			}
			if s.Start.After(f.Transactions[i].Date) {
				break
			}
		}
		if lastTrans.Amount == 0 && lastTrans.Purpose == "" {
			lastTrans.BasicTransaction = s.BasicTransaction
			lastTrans.Date = s.Start
			lastTrans.Status = T_PAID
			ntList = append(ntList, lastTrans)
		}
		//Todo add dates between initial and today
		nd := NextDate(lastTrans.Date, s.Delay, s.DelayType)
		for now.After(nd) {
			ntList = append(ntList, Transaction{
				BasicTransaction: s.BasicTransaction,
				Date:             nd,
				Status:           T_PAID,
			})
			nd = NextDate(nd, s.Delay, s.DelayType)
		}

	}

	f.Transactions = append(f.Transactions, ntList...)
	sort.Sort(Transortable(f.Transactions))

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

type Accumulation struct {
	Transaction
	After int
}

func (f *Family) AccumulateTransactions(ak ACKey) []Accumulation {
	res := []Accumulation{}
	running := 0

	for _, t := range f.Transactions {
		if t.From == ak {
			running -= t.Amount
			res = append(res, Accumulation{t, running})
		}
		if t.Dest == ak {
			running += t.Amount
			res = append(res, Accumulation{t, running})
		}
	}
	return res
}
