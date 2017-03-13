package history

import (
	"fmt"
	"time"
)

func CreateAccount(u, a string) *Account {
	return &Account{
		ACKey:  ACKey{u, a},
		Opened: time.Now(),
	}
}

func (a Account) Merge(b Account) (Account, error) {
	if a.End != b.Start {
		return a, fmt.Errorf("a.End and b.Start not equal")
	}
	if a.ACKey != b.ACKey {
		return a, fmt.Errorf("Merge Account keys don't match")
	}
	return Account{
		ACKey:  a.ACKey,
		Start:  a.Start,
		End:    b.End,
		Opened: a.Opened,
		Closed: b.Closed,
	}, nil
}

func MergeAccountLists(aas, bbs []Account) ([]Account, error) {
	//Loop a, on pair, add on no pair, add

	//Loop b, on no pair add.
	return []Account{}, fmt.Errorf("Merge acounts not ready yet")
}
