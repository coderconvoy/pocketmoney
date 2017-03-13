package history

import "fmt"

func (a Account) Merge(b Account) (Account, error) {
	if a.End != b.Start {
		return a, fmt.Errorf("a.End and b.Start not equal")
	}
	if a.Id != b.Id {
		return a, fmt.Errorf("Merge Account Id's don't match")
	}
	return Account{
		Id:     a.Id,
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
