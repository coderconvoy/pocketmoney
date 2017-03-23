package main

import (
	"math/rand"

	"github.com/coderconvoy/pocketmoney/history"
)

func (f *Family) Calculate() bool {
	res := false
	for _, o := range f.Standing {
		nt, ok := o.Next()
		for ok {
			res = true
			f.ApplyTransaction(nt)
			nt, ok = o.Next()
		}
	}
	return res
}

func (f *Family) AccumulateTransactions(ak history.ACKey) []history.Accumulation {
	return f.Period.Accumulate(ak)

}

func (f *Family) NewStandingID() int32 {
	for {
		n := rand.Int31n(1000)
		fnd := false
		for _, s := range f.Standing {
			if s.ID == n {
				fnd = true
			}
		}
		if !fnd {
			return n
		}
	}
}

func (f *Family) NewRequestID() int32 {
	for {
		n := rand.Int31n(1000)
		fnd := false
		for _, s := range f.Requests {
			if s.ID == n {
				fnd = true
			}
		}
		if !fnd {
			return n
		}
	}
}
