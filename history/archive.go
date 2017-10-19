package history

import (
	"time"

	"github.com/coderconvoy/money"
)

type arcMember struct {
	ACKey
	m money.M
}

type Archive struct {
	start        time.Time
	end          time.Time
	sum          []arcMember
	transactions []Transaction
}

func NewArchive(tt []Transaction) *Archive {
	res := &Archive{transactions: tt}
	for _, v := range tt {
		res.add(v.From, -v.Amount)
		res.add(v.Dest, v.Amount)
	}
	return res
}

func (a *Archive) add(akey ACKey, m money.M) {
	for k, v := range a.sum {
		if v.ACKey == akey {
			a.sum[k].m += m
			return
		}
	}
	a.sum = append(a.sum, arcMember{ACKey: akey, m: m})
}
