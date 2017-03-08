package main

import (
	"testing"
	"time"
)

func Test_Standing(t *testing.T) {
	fam := Family{
		Transactions: []Transaction{
			qtrans("a", "b", "C", "C", 10, "s1", 2017, 3, 3),
		},
	}

}

func qstanding(fu, du, fa, da string, n int, purp string, sy, sm, sd, d, dt int) StandingOrder {
	return StandingOrder{
		BasicTransaction: BasicTransaction{fu, du, fa, da, n, purp},
		Date:             qdate(y, m, d),
		Delay:            d,
		DelayType:        dt,
	}
}

func qtrans(fu, du, fa, da string, n int, purp string, y, m, d int) Transaction {
	return Transaction{
		BasicTransaction: BasicTransaction{fu, du, fa, da, n, purp},
		Date:             qdate(y, m, d),
		Status:           T_PAID,
	}

}

func qdate(y, m, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}
