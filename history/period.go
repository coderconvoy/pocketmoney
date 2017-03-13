package history

import (
	"fmt"
	"sort"
	"time"
)

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

func (p Period) Split(tt ...time.Time) []Period {
	cp := Period{}
	for _, v := range p.Accounts {
		na := v
		na.End = na.Start
		cp.Accounts = append(cp.Accounts, na)
	}

	nsplit := 0
	res := []Period{}

	sort.Sort(Transortable(p.Transactions))

	for _, tac := range p.Transactions {
		//Move to next split, normally only once
		for nsplit < len(tt) && tac.Date.After(tt[nsplit]) {
			res = append(res, cp)
			cp = cp.StartNext()
			nsplit++
		}

		//find from and to accounts and apply transaction
		var fr, ds *Account

		for i, ac := range cp.Accounts {
			if ac.Id == tac.From {
				fr = &cp.Accounts[i]
			}
			if ac.Id == tac.Dest {
				ds = &cp.Accounts[i]
			}
		}
		if !(fr == nil || ds == nil) {
			fr.End -= tac.Amount
			ds.End += tac.Amount

		}
		cp.Transactions = append(cp.Transactions, tac)

	}
	res = append(res, cp)
	return res
}

func (p Period) Merge(p2 Period) (Period, error) {
	return p, fmt.Errorf("Merge not ready yet")
}

func (p Period) Accumulate(ak ACKey) []Accumulation {
	running := 0
	for _, ac := range p.Accounts {
		if ak == ac.Id {
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
