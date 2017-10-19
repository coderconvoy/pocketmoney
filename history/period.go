package history

import (
	"fmt"
	"sort"
	"time"

	"github.com/coderconvoy/dbase"
	"github.com/coderconvoy/money"
)

type Period struct {
	Start, End   time.Time
	Accounts     []*Account
	Transactions []Transaction
}

func (p Period) StartNext() Period {
	res := Period{}
	for _, v := range p.Accounts {
		nx := v
		nx.Start = v.End
		nx.End = v.End
		res.Accounts = append(res.Accounts, nx)
	}
	return res
}

func (p *Period) ApplyTransaction(ts ...Transaction) error {
	dbase.QLog("Applying transaction")
	fmt.Println("Applying ftransaction")
	if len(ts) == 0 {
		return nil
	}
	dbase.QLog(fmt.Sprintln(ts[0]))
	for _, t := range ts {
		if p.Start.After(t.Date) || (p.End.After(time.Time{}) && t.Date.After(p.End)) {
			return fmt.Errorf("All dates must fit within range")
		}

		ffnd, dfnd := false, false
		for _, a := range p.Accounts {
			if a.ACKey == t.From {
				ffnd = true
			}
			if a.ACKey == t.Dest {
				dfnd = true
			}
		}
		if !ffnd || !dfnd {
			return fmt.Errorf("Transaction does not have matching account")
		}
	}
	p.Transactions = append(p.Transactions, ts...)

	sort.Sort(Transortable(p.Transactions))

	for _, a := range p.Accounts {
		a.End = a.Start
	}
	for _, t := range p.Transactions {
		for _, a := range p.Accounts {
			if a.ACKey == t.From {
				a.End -= t.Amount
			}
			if a.ACKey == t.Dest {
				a.End += t.Amount
			}
		}
	}
	return nil
}

func (p Period) Accumulate(ak ACKey) []Accumulation {
	running := money.M(0)
	for _, ac := range p.Accounts {
		if ak == ac.ACKey {
			running = ac.Start
			break
		}
	}
	res := []Accumulation{}

	for _, t := range p.Transactions {
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

func (p Period) UserAccounts(uname string) []*Account {
	res := []*Account{}
	for _, a := range p.Accounts {
		if a.Username == uname {
			res = append(res, a)
		}
	}
	return res
}
