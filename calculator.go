package main

import (
	"errors"
	"fmt"
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
