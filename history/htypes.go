package history

import (
	"errors"
	"fmt"
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

func ParseAmount(st string) (int, error) {
	dotted := -1
	res := 0
	for _, s := range st {
		if dotted >= 0 {
			dotted++
			if dotted > 3 {
				return 0, errors.New("dp too small")
			}
		}
		switch s {
		case '.':
			if dotted >= 0 {
				return 0, errors.New("Too many dots")
			}
			dotted = 1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n := int(s - '0')
			res = res*10 + n
		default:
			return 0, errors.New("Unexpected Char: " + string(s))
		}

	}
	if dotted == -1 {
		return res * 100, nil
	}
	for dotted < 3 {
		res *= 10
		dotted++
	}
	return res, nil

	/*am, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.New("Could not parse amount")
	}
	if am < 0 {
		return 0, errors.New("Amount not Positive")
	}
	return int(am * 100), nil*/
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
