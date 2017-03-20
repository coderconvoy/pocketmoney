package history

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ACKey struct {
	Username, Name string
}

func NewACKey(s string) (ACKey, error) {
	sp := strings.Split(s, ":")
	if len(sp) != 2 {
		return ACKey{}, fmt.Errorf("Could not Parse '%s'.", s)
	}
	return ACKey{sp[0], sp[1]}, nil
}

func ParseAmount(s string) (int, error) {
	am, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.New("Could not parse amount")
	}
	if am < 0 {
		return 0, errors.New("Amount not Positive")
	}
	return int(am * 100), nil
}

func (a ACKey) String() string { return a.Username + ":" + a.Name }

type Accumulation struct {
	Transaction
	After int
}
type Account struct {
	ACKey
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
	Accounts     []*Account
	Transactions []Transaction
}
