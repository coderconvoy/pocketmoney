package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	dbase "github.com/coderconvoy/dbase2"
)

var GT *template.Template
var UserDB = dbase.DBase{"data/users"}

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, "--", r.URL.Path)
	GT.ExecuteTemplate(w, "index.html", nil)
}

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	var u User
	u.FName = r.FormValue("Full Name")
	u.UName = r.FormValue("Username")
	d, err := json.Marshal(u)
	if err != nil {
		fmt.Fprintf(w, "No Can Add: %s", err)
		return
	}
	UserDB.WriteMap(u.UName, d)
	http.SetCookie(w, &http.Cookie{
		Name:  "Username",
		Value: u.UName,
	})
	GT.ExecuteTemplate(w, "newuser.html", u)
}

func HandleNewAccount(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Username")
	if err != nil {
		GT.ExecuteTemplate(w, "index.html", nil)
	}
	b, err := UserDB.ReadMap(c.Value)
	if err != nil {
		fmt.Fprintf(w, "No thing stored for user %s", c.Value)
	}

	GT.ExecuteTemplate(w, "newaccount.html", b)

}

func main() {
	/*t, err := Asset("assets/index.html")
	if err != nil {
		fmt.Println("ERRor:", err)
		return
	}*/
	GT = template.New("index")
	ad, err := AssetDir("assets")
	for _, n := range ad {
		t, err := Asset("assets/" + n)
		fmt.Println("Parsing :" + n)
		GT = GT.New(n)
		_, err = GT.Parse(string(t))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	http.HandleFunc("/", Handle)
	http.HandleFunc("/NewUser", HandleNewUser)
	http.HandleFunc("/NewAccount", HandleNewAccount)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
