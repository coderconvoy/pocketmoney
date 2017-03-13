package history

import "time"

type ACKey struct {
	Username, Name string
}
type Account struct {
	Id             ACKey
	Opened, Closed time.Time
	Start, End     int
}

type Transaction struct {
	From, Dest ACKey
	Amount     int
	Purpose    string
	Date       time.Time
}

type Transortable []Transaction

func (t Transortable) Len() int           { return len(t) }
func (t Transortable) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Transortable) Less(i, j int) bool { return t[j].Date.After(t[i].Date) }

type Period struct {
	Start, End   time.Time
	Accounts     []Account
	Transactions []Transaction
}
