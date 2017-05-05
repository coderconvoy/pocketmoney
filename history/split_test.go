package history

import (
	"fmt"
	"testing"
	"time"
)

func Test_split(t *testing.T) {
	p := Period{
		Start: qdate(2017, 01, 01),
		End:   qdate(2017, 03, 01),
		Accounts: []*Account{
			qac(AC1, 0, 0),
			qac(AC2, 5, 0),
			qac(AC3, 0, 0),
		},
		Transactions: []Transaction{
			qtran(AC1, AC2, 50, "a", qdate(2017, 01, 01)),
			qtran(AC2, AC3, 50, "a", qdate(2017, 01, 02)),
			qtran(AC1, AC2, 50, "a", qdate(2017, 01, 03)),
			qtran(AC1, AC2, 50, "a", qdate(2017, 01, 04)),
			qtran(AC1, AC2, 50, "a", qdate(2017, 01, 05)),
		},
	}
	sp := p.Split(qdate(2017, 01, 03))
	if len(sp) != 2 {
		t.Logf("Not 2 responses")
		t.Fail()
	}

	for r, p := range sp {
		fmt.Println("Period:", r)
		fmt.Println(" AC:")
		for _, ac := range p.Accounts {
			fmt.Println("  ", ac.Username, ac.Start, ac.End)
		}
		fmt.Println(" TC:")
		for _, tr := range p.Transactions {
			fmt.Println("  ", tr.From, tr.Dest)
		}
	}

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

func qtran(f, d, a int, p string, t time.Time) Transaction {
	return Transaction{
		From:    qkey(f),
		Dest:    qkey(d),
		Amount:  a,
		Purpose: p,
		Date:    t,
	}
}

func qac(k, s, e int) *Account {
	return &Account{
		ACKey: qkey(k),
		Start: s,
		End:   e,
	}
}
func qdate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}
