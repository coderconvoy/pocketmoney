package main

import (
	"net/http"

	"github.com/coderconvoy/dbase2"
)

func HandlePasswordChange(w http.ResponseWriter, r *http.Request) {
	fam, fmem, err := LoggedInFamily(w, r)
	if err != nil {
		GoIndex(w, r, "Not Logged In"+err.Error())
		return
	}

	var cmem *User
	for i, m := range fam.Members {
		if m.Username == fmem {
			cmem = &fam.Members[i]
		}
	}

	oldpwd := r.FormValue("oldpwd")
	if !cmem.Password.Check(oldpwd) {
		ExTemplate(GT, w, "userhome.html", PageData{"Error Saving password : Old password incorrect", fmem, fam})
		return
	}

	pwd1 := r.FormValue("pwd1")
	pwd2 := r.FormValue("pwd2")
	if len(pwd1) < 5 {
		ExTemplate(GT, w, "userhome.html", PageData{"Error Saving password : New password too short (5 min)", fmem, fam})
		return
	}

	if pwd1 != pwd2 {
		ExTemplate(GT, w, "userhome.html", PageData{"Error Saving password : Password confirmation doesn't match", fmem, fam})
		return
	}

	np, err := dbase2.NewPassword(pwd1)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", PageData{"Error Saving password : " + err.Error(), fmem, fam})
		return
	}

	cmem.Password = np

	err = SaveFamily(fam)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", PageData{"Error Saving password : " + err.Error(), fmem, fam})
		return
	}

	ExTemplate(GT, w, "userhome.html", PageData{"Saved", fmem, fam})

}
