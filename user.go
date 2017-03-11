package main

import (
	"github.com/coderconvoy/dbase2"
)

func HandlePasswordChange(ld LoginData) {
	w, r, fam, fmem := ld.W, ld.R, ld.Fam, ld.Fmem

	var cmem *User
	for i, m := range fam.Members {
		if m.Username == fmem {
			cmem = &fam.Members[i]
		}
	}

	oldpwd := r.FormValue("oldpwd")
	if !cmem.Password.Check(oldpwd) {
		ExTemplate(GT, w, "userhome.html", ld.Pd("Error Saving password : Old password incorrect"))
		return
	}

	pwd1 := r.FormValue("pwd1")
	pwd2 := r.FormValue("pwd2")
	if len(pwd1) < 5 {
		ExTemplate(GT, w, "userhome.html", ld.Pd("Error Saving password : New password too short (5 min)"))
		return
	}

	if pwd1 != pwd2 {
		ExTemplate(GT, w, "userhome.html", ld.Pd("Error Saving password : Password confirmation doesn't match"))
		return
	}

	np, err := dbase2.NewPassword(pwd1)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd("Error Saving password : "+err.Error()))
		return
	}

	cmem.Password = np

	ExTemplate(GT, w, "userhome.html", ld.Pd(""))
}
