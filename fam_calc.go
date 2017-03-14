package main

import (
	"github.com/coderconvoy/pocketmoney/history"
)

func (f *Family) Calculate() bool {
	res := false
	for _, o := range f.Standing {
		nt, ok := o.Next()
		for ok {
			res = true
			f.Period.ApplyTransaction(nt)
			nt, ok = o.Next()
		}
	}
	return res

}

func (f *Family) AccumulateTransactions(ak history.ACKey) []history.Accumulation {
	return f.Period.Accumulate(ak)

}
