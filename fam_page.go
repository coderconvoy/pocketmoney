package main

import (
	"fmt"
	"time"

	"github.com/coderconvoy/dbase2"
)

func HandleFamily(ld LoginData) {
	fmt.Println("Going Family")
	ExTemplate(GT, ld.W, "familypage.html", NewPageData("", ld.Fmem, ld.Fam))
}
func HandleAddMember(ld LoginData) {
	w, r, fam, fmem := ld.W, ld.R, ld.Fam, ld.Fmem
	if !IsParent(fmem, fam) {
		ExTemplate(GT, w, "familypage.html", NewPageData("Not a Parent", fmem, fam))
		return
	}

	uname := r.FormValue("username")
	parent := r.FormValue("parent")
	fmt.Println("Parent == " + parent)
	pwd1 := r.FormValue("pwd1")
	pwd2 := r.FormValue("pwd2")

	if uname == "" {
		ExTemplate(GT, w, "familypage.html", NewPageData("No Username", fmem, fam))
		return
	}
	if pwd1 != pwd2 || pwd1 == "" {
		ExTemplate(GT, w, "familypage.html", NewPageData("Passwords not matching", fmem, fam))
		return
	}

	pw, err := dbase2.NewPassword(pwd1)
	if err != nil {
		fmt.Println("Could not Password: ", err)
		GoIndex(w, r, "Password Failed")
		return
	}
	fam.Members = append(fam.Members, User{
		Username: uname,
		Parent:   parent == "on",
		Password: pw,
	})
	fam.Accounts = append(fam.Accounts, &Account{
		ACKey:     ACKey{uname, "checking"},
		Current:   0,
		StartDate: time.Now(),
	})

	err = fam.Save()
	if err != nil {
		//TODO
	}
	ExTemplate(GT, w, "familypage.html", NewPageData("", fmem, fam))
}
