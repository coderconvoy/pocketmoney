package main

import (
	"testing"
	"time"

	"github.com/coderconvoy/pocketmoney/history"
)

func Test_Standing(t *testing.T) {
	_ = &Family{
		Period: history.Period{
			Transactions: []history.Transaction{
				qtrans("a", "b", "C", "C", 20, "s1", qdate(2017, 3, 3)),
				qtrans("a", "f", "C", "C", 400, "sev", qdate(2017, 3, 5)),
			},
		},
		Standing: []*StandingOrder{
			qstanding("a", "b", "C", "C", 20, "s1", qdate(2017, 2, 3), 7, D_NDAYS), //expects 2
			qstanding("a", "b", "C", "C", 20, "s2", qdate(2017, 2, 3), 7, D_NDAYS), //expects 5
		},
	}

	/*	fam.Calculate(qdate(2017, 3, 15))
		if len(fam.Transactions) != 9 {
			t.Logf("fam Transactions len, ex : 9, got : %d", len(fam.Transactions))
			for _, tr := range fam.Transactions {
				t.Logf("%s", tr)
			}
			t.Fail()
		}
	*/

}

func qstanding(fu, du, fa, da string, n int, purp string, dat time.Time, d, dt int) *StandingOrder {
	return &StandingOrder{
		Transaction:  history.Transaction{history.ACKey{fu, fa}, history.ACKey{du, da}, n, purp, time.Time{}},
		Start:        dat,
		Interval:     d,
		IntervalType: dt,
	}
}

func qtrans(fu, du, fa, da string, n int, purp string, dt time.Time) history.Transaction {
	return history.Transaction{history.ACKey{fu, fa}, history.ACKey{du, da}, n, purp, dt}

}

func qdate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}
