package main

import (
	"testing"
	"time"
)

func Test_Standing(t *testing.T) {
	fam := &Family{
		Transactions: []Transaction{
			qtrans("a", "b", "C", "C", 20, "s1", qdate(2017, 3, 3)),
			qtrans("a", "f", "C", "C", 400, "sev", qdate(2017, 3, 5)),
		},
		Standing: []StandingOrder{
			qstanding("a", "b", "C", "C", 20, "s1", qdate(2017, 2, 3), 7, D_NDAYS), //expects 2
			qstanding("a", "b", "C", "C", 20, "s2", qdate(2017, 2, 3), 7, D_NDAYS), //expects 5
		},
	}

	fam.calculateStanding(qdate(2017, 3, 15))
	if len(fam.Transactions) != 9 {
		t.Logf("fam Transactions len, ex : 9, got : %d", len(fam.Transactions))
		for _, tr := range fam.Transactions {
			t.Logf("%s", tr)
		}
		t.Fail()
	}

}

func qstanding(fu, du, fa, da string, n int, purp string, dat time.Time, d, dt int) StandingOrder {
	return StandingOrder{
		BasicTransaction: BasicTransaction{fu, du, fa, da, n, purp},
		Start:            dat,
		Delay:            d,
		DelayType:        dt,
	}
}

func qtrans(fu, du, fa, da string, n int, purp string, dt time.Time) Transaction {
	return Transaction{
		BasicTransaction: BasicTransaction{fu, du, fa, da, n, purp},
		Date:             dt,
		Status:           T_PAID,
	}

}

func qdate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}
