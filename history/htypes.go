package history

import (
	"fmt"
	"strings"
	"time"

	"github.com/coderconvoy/money"
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

func (a ACKey) String() string { return a.Username + ":" + a.Name }

type Accumulation struct {
	Transaction
	After money.M
}
type Account struct {
	ACKey
	Opened, Closed time.Time
	Start, End     money.M
	Col1, Col2     string
}

type Transaction struct {
	From, Dest ACKey
	Amount     money.M
	Purpose    string
	Date       time.Time
}

type Transortable []Transaction

func (t Transortable) Len() int           { return len(t) }
func (t Transortable) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Transortable) Less(i, j int) bool { return t[j].Date.After(t[i].Date) }
