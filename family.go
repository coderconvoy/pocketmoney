package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/coderconvoy/dbase2"
	"github.com/pkg/errors"
)

// LoadFamily Reads and Unmarshals the Family File
// Params family name,
// Returns loaded family or nil, followed by error
func LoadFamily(fname string) (*Family, error) {
	d, err := FamDB.ReadMap(fname)
	if err != nil {
		return nil, errors.Wrap(err, "No Family Exists")
	}

	var f Family
	err = json.Unmarshal(d, &f)
	if err != nil {
		return nil, errors.Wrap(err, "Corrupted Family File")
	}
	return &f, nil
}

func SaveFamily(f *Family) error {
	mar, err := json.MarshalIndent(&f, "", " ")
	if err != nil {
		return err
	}

	ok := FamDB.WriteMap(f.FamilyName, mar)
	if !ok {
		return errors.New("Could not save Family")
	}
	return nil
}

func HandleFamily(ld LoginData) {
	fmt.Println("Going Family")
	ExTemplate(GT, ld.W, "familypage.html", PageData{"", ld.Fmem, ld.Fam})
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	//Load Family File
	famname := r.FormValue("family")
	fam, err := LoadFamily(famname)
	if err != nil {
		GoIndex(w, r, err.Error())
		return
	}

	uname := r.FormValue("username")

	http.SetCookie(w, &http.Cookie{
		Name:    "LastLog",
		Value:   famname + ":" + uname,
		Expires: time.Now().Add(time.Hour * 24 * 30),
	})

	//Check Password
	pw := r.FormValue("pwd")

	sel := -1
	for i, v := range fam.Members {
		if v.Username == uname || (v.Email == uname) {
			if v.Password.Check(pw) {
				sel = i
				break
			}
		}
	}
	if sel == -1 {
		ExTemplate(GT, w, "index.html", IndexData{"No Username-Password match", famname, uname})
		return
	}

	loginControl.Login(w, fam.FamilyName, uname)
	ExTemplate(GT, w, "familypage.html", PageData{"", uname, fam})

}
func HandleNewFamily(w http.ResponseWriter, r *http.Request) {
	var f Family
	f.FamilyName = r.FormValue("familyname")
	uname := r.FormValue("username")

	http.SetCookie(w, &http.Cookie{
		Name:    "LastLog",
		Value:   f.FamilyName + ":" + uname,
		Expires: time.Now().Add(time.Hour * 24 * 30),
	})

	//Check if family already exists
	_, err := FamDB.ReadMap(f.FamilyName)
	if err == nil {
		GoIndex(w, r, "Family Name Already Exists")
		return
	}

	p1 := r.FormValue("pwd1")
	p2 := r.FormValue("pwd2")

	if p1 != p2 {
		GoIndex(w, r, "Passwords don't match")
		return
	}

	pw, err := dbase2.NewPassword(p1)
	if err != nil {
		GoIndex(w, r, "Passwords error: "+err.Error())
		return
	}

	email := r.FormValue("email")
	f.Members = append(f.Members, User{
		Username: uname,
		Email:    email,
		Password: pw,
		Parent:   true,
	})

	f.Accounts = append(f.Accounts,
		&Account{
			Username:  "WORLD",
			Name:      "main",
			StartDate: time.Now(),
			Current:   0,
		},
		&Account{
			Username:  uname,
			Name:      "checking",
			StartDate: time.Now(),
			Current:   0,
		})

	err = SaveFamily(&f)
	if err != nil {
		//TODO
	}
	loginControl.Login(w, f.FamilyName, uname)
	ExTemplate(GT, w, "familypage.html", PageData{"", uname, &f})

}

func HandleAddMember(ld LoginData) {
	w, r, fam, fmem := ld.W, ld.R, ld.Fam, ld.Fmem
	if !IsParent(fmem, fam) {
		ExTemplate(GT, w, "familypage.html", PageData{"Not a Parent", fmem, fam})
		return
	}

	uname := r.FormValue("username")
	parent := r.FormValue("parent")
	fmt.Println("Parent == " + parent)
	pwd1 := r.FormValue("pwd1")
	pwd2 := r.FormValue("pwd2")

	if uname == "" {
		ExTemplate(GT, w, "familypage.html", PageData{"No Username", fmem, fam})
		return
	}
	if pwd1 != pwd2 || pwd1 == "" {
		ExTemplate(GT, w, "familypage.html", PageData{"Passwords not matching", fmem, fam})
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
		Username:  uname,
		Name:      "Checking",
		Current:   0,
		StartDate: time.Now(),
	})

	err = SaveFamily(fam)
	if err != nil {
		//TODO
	}
	ExTemplate(GT, w, "familypage.html", PageData{"", fmem, fam})
}
