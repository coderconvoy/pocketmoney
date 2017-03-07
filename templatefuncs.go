package main

import (
	"errors"
	"text/template"
)

func TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"getuser":  GetUser,
		"isparent": IsParent,
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

func IsParent(uname string, fam *Family) bool {
	m, err := GetUser(uname, fam)
	if err != nil {
		return false
	}
	return m.Parent
}
