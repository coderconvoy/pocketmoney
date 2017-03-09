package main

import (
	"errors"
	"fmt"
	"text/template"
	"time"
)

func TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"getuser":  GetUser,
		"isparent": IsParent,
		"money":    PrintMoney,
		"date":     PrintDate,
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

func PrintDate(t time.Time) string {
	return t.Format("Mon 2/Jan/06")
}

func IsParent(uname string, fam *Family) bool {
	m, err := GetUser(uname, fam)
	if err != nil {
		return false
	}
	return m.Parent
}
