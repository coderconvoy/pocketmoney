package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func (f *Family) Calculate() {
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

	var from, dest *Account = nil, nil

	for _, a := range f.Accounts {
		if a.Username == t.FromUser && a.Name == t.FromAC {
			from = a
		}
		if a.Username == t.DestUser && a.Name == t.DestAC {
			dest = a
		}
	}
	if from == nil || dest == nil {
		return errors.New("No Account for Transaction")
	}
	from.Current -= t.Amount
	dest.Current += t.Amount
	return nil
}

func (f *Family) CalculateStanding() {
	f.CalculateStanding(time.Now())
}

func (f *Family) calculateStanding(now time.Time) {
	ntList := []Transaction{}
	for _, s := range f.Standing {

		//Some standing orders will be set in the future.
		if s.Start.After(time.Now()) {
			continue
		}

		var lastTrans Transaction

		for i := len(f.Transactions) - 1; i >= 0; i-- {
			if f.Transactions[i].Purpose == s.Purpose {
				lastTrans = f.Transactions[i]
				break
			}
			if s.Start.After(t.Transactions[i].Date) {
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
		for time.Now().After(nd) {
			ntList = append(ntList, Transaction{
				BasicTransaction: s.BasicTransaction,
				Date:             nd,
				Status:           T_PAID,
			})
			nd := NextDate(nd, s.Delay, s.DelayType)
		}

	}

	fam.Transactions = append(fam.Transactions, ntList...)
	sort.Sort(Transortable(fam.Transactions))

}

func NextDate(d time.Time, step int, steptype int) time.Time {
	if steptype == D_NDAYS {
		return d.AddDate(0, 0, step)
	}
	return d.AddDate(0, step, 0)
}
