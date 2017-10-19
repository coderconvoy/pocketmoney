package history

import (
	"fmt"
	"testing"
	"time"

	"github.com/coderconvoy/money"
)

func Test_Create(t *testing.T) {
	transactions := []Transaction{
		qtran(AC1, AC2, 50, "a", qdate(2017, 01, 01)),
		qtran(AC2, AC3, 50, "a", qdate(2017, 01, 02)),
		qtran(AC1, AC2, 50, "a", qdate(2017, 01, 03)),
		qtran(AC1, AC2, 50, "a", qdate(2017, 01, 04)),
		qtran(AC1, AC2, 50, "a", qdate(2017, 01, 05)),
	}

	a := NewArchive(transactions)

	tot := money.M(0)
	for _, v := range a.sum {
		fmt.Println(v.m)
		tot += v.m
	}
	if tot != 0 {
		t.Errorf("Archive should sum to Zero, got: %d", tot)
	}
	fmt.Println(a.sum)

}

const (
	AC1 = iota
	AC2
	AC3
	AC4
)

func qkey(n int) ACKey {
	switch n {
	case AC1:
		return ACKey{"a", "ch"}
	case AC2:
		return ACKey{"b", "ch"}
	case AC3:
		return ACKey{"c", "ch"}
	}
	return ACKey{"d", "ch"}
}
func qtran(f, d int, a money.M, p string, t time.Time) Transaction {
	return Transaction{
		From:    qkey(f),
		Dest:    qkey(d),
		Amount:  a,
		Purpose: p,
		Date:    t,
	}
}

func qac(k int, s, e money.M) *Account {
	return &Account{
		ACKey: qkey(k),
		Start: s,
		End:   e,
	}
}

func qdate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}
