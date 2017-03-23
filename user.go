package main

import (
	"github.com/coderconvoy/dbase2"
)

func HandlePasswordChange(ld *PageHand) (string, string) {
	r, fam, fmem := ld.R, ld.Fam, ld.Fmem

	var cmem *User
	for i, m := range fam.Members {
		if m.Username == fmem {
			cmem = &fam.Members[i]
		}
	}

	oldpwd := r.FormValue("oldpwd")
	if !cmem.Password.Check(oldpwd) {
		return "/personal", "Old Password Incorrect"
	}

	pwd1 := r.FormValue("pwd1")
	pwd2 := r.FormValue("pwd2")
	if len(pwd1) < 5 {
		return "/personal", "New Password too short (5 chars min)"
	}

	if pwd1 != pwd2 {
		return "/personal", "Password Confirmation doesn't match"
	}

	np, err := dbase2.NewPassword(pwd1)
	if err != nil {
		return "/personal", "Password error: " + err.Error()
	}

	cmem.Password = np

	return "/.personal", ""
}
