package main

import (
	"time"

	"github.com/coderconvoy/dbase2"
	"github.com/coderconvoy/pocketmoney/history"
)

// HandleFamily show the Family page
func HandleFamily(ld LoginData) {
	ExTemplate(GT, ld.W, "familypage.html", NewPageData("", ld.Fmem, ld.Fam))
}

func HandleAddMember(ld LoginData) {
	w, r, fam, fmem := ld.W, ld.R, ld.Fam, ld.Fmem
	if !fam.IsParent(fmem) {
		ExTemplate(GT, w, "familypage.html", NewPageData("Not a Parent", fmem, fam))
		return
	}

	uname := r.FormValue("username")
	parent := r.FormValue("parent")
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
		dbase2.QLog("Could not Password: " + err.Error())
		GoIndex(w, r, "Password Failed")
		return
	}
	fam.Members = append(fam.Members, User{
		Username: uname,
		Parent:   parent == "on",
		Password: pw,
	})
	fam.Period.Accounts = append(fam.Period.Accounts, &history.Account{
		ACKey:  history.ACKey{uname, "checking"},
		Opened: time.Now(),
	})

	ExTemplate(GT, w, "familypage.html", NewPageData("", fmem, fam))
}
