package main

import (
	"errors"
	"fmt"
	"text/template"
	"time"
)

func TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"getuser":        GetUser,
		"isparent":       IsParent,
		"accessACs":      FilterWriteableACs,
		"standingbyac":   FilterStandingByAC,
		"standingbyuser": FilterStandingByUser,
		"money":          PrintMoney,
		"date":           PrintDate,
		"dateRFC":        PrintDateRFC,
	}
}

func GetUser(uname string, fam *Family) (*User, error) {
	for i, m := range fam.Members {
		if m.Username == uname {
			return &fam.Members[i], nil
		}
	}
	return nil, errors.New("No Member of that name")
}

func PrintMoney(n int) string {
	if n < 0 {
		return "-£" + fmt.Sprintf("%.2f", float32(-n)/100)
	}
	return "£" + fmt.Sprintf("%.2f", float32(n)/100)
}

func PrintDate(t ...time.Time) string {
	if len(t) == 0 {
		return time.Now().Format("Mon 2/Jan/06")
	}
	return t[0].Format("Mon 2/Jan/06")
}

func PrintDateRFC(t ...time.Time) string {
	if len(t) == 0 {
		return time.Now().Format("2006-01-02")
	}
	return t[0].Format("2006-01-02")
}

func IsParent(uname string, fam *Family) bool {
	m, err := GetUser(uname, fam)
	if err != nil {
		return false
	}
	return m.Parent
}

func WriteableAC(a *Account, uname string, fam *Family) bool {
	isPar := IsParent(uname, fam)
	return a.Username == uname || (a.Username == "WORLD" && isPar)
}

func FilterWriteableACs(acs []*Account, uname string, fam *Family) []*Account {
	res := []*Account{}
	for _, v := range acs {
		if WriteableAC(v, uname, fam) {
			res = append(res, v)
		}
	}
	return res
}

func FilterStandingByUser(st []StandingOrder, uname string) []StandingOrder {
	res := []StandingOrder{}
	for _, v := range st {
		if (uname == v.FromUser) || (uname == v.DestUser) {

			res = append(res, v)
		}
	}
	return res
}

func FilterStandingByAC(st []StandingOrder, ac *Account) []StandingOrder {
	res := []StandingOrder{}
	for _, v := range st {
		if (ac.Name == v.FromAC && ac.Username == v.FromUser) ||
			(ac.Name == v.DestAC && ac.Username == v.DestUser) {

			res = append(res, v)
		}
	}
	return res
}
