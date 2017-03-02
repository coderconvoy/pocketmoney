package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coderconvoy/dbase2"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	//Load Family File
	e := r.FormValue("family")
	var f Family
	d, err := FamDB.ReadMap(e)
	if err != nil {
		ExTemplate(GT, w, "index.html", "No Family Exists")
		return
	}
	err = json.Unmarshal(d, &f)
	if err != nil {
		ExTemplate(GT, w, "index.html", "Corrupted Family File")
		return
	}

	//Check Password
	uname := r.FormValue("username")
	pw := r.FormValue("pwd")
	if uname == "" {
		ExTemplate(GT, w, "index.html", "Username Blank")
	}

	sel := -1
	for i, v := range f.Members {
		if v.Username == uname || (v.Email == uname) {
			if v.Password.Check(pw) {
				sel = i
				break
			}
		}
	}
	if sel == -1 {
		ExTemplate(GT, w, "index.html", "No Username-Password match")
		return
	}

	loginControl.Login(w, f.FamilyName, uname)
	ExTemplate(GT, w, "familypage.html", f)

}
func HandleNewFamily(w http.ResponseWriter, r *http.Request) {
	var f Family
	f.FamilyName = r.FormValue("familyname")

	//Check if family already exists
	_, err := FamDB.ReadMap(f.FamilyName)
	if err == nil {
		ExTemplate(GT, w, "index.html", "Email Already Exists")
		return
	}

	p1 := r.FormValue("pwd1")
	p2 := r.FormValue("pwd2")

	if p1 != p2 {
		ExTemplate(GT, w, "index.html", "Passwords don't match")
		return
	}

	pw, err := dbase2.NewPassword(p1)
	if err != nil {
		ExTemplate(GT, w, "index.html", "Passwords error: "+err.Error())
		return
	}

	uname := r.FormValue("username")
	email := r.FormValue("email")
	f.Members = append(f.Members, User{
		Username: uname,
		Email:    email,
		Password: pw,
		Parent:   true,
	})

	mar, err := json.Marshal(&f)
	if err != nil {
		fmt.Println("could not marshal f:", err)
	}

	FamDB.WriteMap(f.FamilyName, mar)
	loginControl.Login(w, f.FamilyName, uname)
	ExTemplate(GT, w, "familypage.html", f)

}

func HandleAddMember(w http.ResponseWriter, r *http.Request) {
	a, ok := loginControl.GetLogin(w, r)
	if ok != dbase2.OK {
		ExTemplate(GT, w, "index.html", "Cannot add member: you are not logged in")
		return
	}
	fmap, err := FamDB.ReadMap(a.Familyname)
	if err != nil {
		ExTemplate(GT, w, "index.html", "Family "+a.Familyname+" Could not be loaded")
		return
	}

	var f Family
	err = json.Unmarshal(fmap, &f)
	if err != nil {
		fmt.Println("could not unmarshal a.FamilyName: " + err.Error())
		ExTemplate(GT, w, "index.html", "Corrupt File")
		return
	}

	uname := r.FormValue("username")
	parent := r.FormValue("parent")
	fmt.Println("Parent == " + parent)
	pwd1 := r.FormValue("pwd1")
	pwd2 := r.FormValue("pwd2")

	if uname == "" {
		ExTemplate(GT, w, "index.html", "No Username")
		return
	}
	if pwd1 != pwd2 || pwd1 == "" {
		ExTemplate(GT, w, "index.html", "Passwords not matching ")
		return
	}

	pw, err := dbase2.NewPassword(pwd1)
	if err != nil {
		fmt.Println("Could not Password: ", err)
		ExTemplate(GT, w, "index.html", "Password Failed")
	}
	f.Members = append(f.Members, User{
		Username: uname,
		Parent:   parent == "true",
		Password: pw,
	})

	mar, err := json.Marshal(f)
	if err != nil {
		ExTemplate(GT, w, "index.html", "Could not Marshal: "+err.Error())
		return
	}
	wrote := FamDB.WriteMap(f.FamilyName, mar)
	if !wrote {
		ExTemplate(GT, w, "index.html", "Could not write new member to db")
		return
	}

	ExTemplate(GT, w, "familypage.html", f)

}
