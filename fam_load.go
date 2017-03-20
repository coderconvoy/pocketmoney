package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/coderconvoy/dbase2"
	"github.com/coderconvoy/pocketmoney/history"
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
	if len(f.Members) == 0 {
		return nil, errors.New("No Family Members")
	}

	if f.Members[0].Username != "WORLD" {
		f.Members = append([]User{
			{
				Username: "WORLD",
			},
		}, f.Members...)
	}
	return &f, nil
}

func (f *Family) Save() error {
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
	ExTemplate(GT, w, "familypage.html", NewPageData("", uname, fam))

}
func HandleNewFamily(w http.ResponseWriter, r *http.Request) {
	f := &Family{}
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

	f.Period = history.Period{
		Accounts: []*history.Account{
			history.CreateAccount("WORLD", "main"),
			history.CreateAccount(uname, "checking"),
		},
	}

	err = f.Save()
	if err != nil {
		//TODO
	}
	loginControl.Login(w, f.FamilyName, uname)
	ExTemplate(GT, w, "familypage.html", NewPageData("", uname, f))

}
