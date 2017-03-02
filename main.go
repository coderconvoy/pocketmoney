package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/coderconvoy/dbase2"
	"github.com/coderconvoy/templater/tempower"
)

var GT *template.Template
var FamDB = dbase2.DBase{"data/families"}

func ExTemplate(t *template.Template, w http.ResponseWriter, name string, data interface{}) {
	err := t.ExecuteTemplate(w, name, data)
	if err != nil {
		fmt.Println(err)
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, "--", r.URL.Path)
	ExTemplate(GT, w, "index.html", nil)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	loginControl.Logout(w, r)
	ExTemplate(GT, w, "index.html", "You are now Logged out")

}

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
}

func HandleNewAccount(w http.ResponseWriter, r *http.Request) {
	/*c, err := r.Cookie("Username")
	if err != nil {
		ExTemplate(GT,w, "index.html", nil)
	}
	b, err := UserDB.ReadMap(c.Value)
	if err != nil {
		fmt.Fprintf(w, "No thing stored for user %s", c.Value)
	}

	ExTemplate(GT,w, "newaccount.html", b)
	*/

}

func main() {
	/*t, err := Asset("assets/index.html")
	if err != nil {
		fmt.Println("ERRor:", err)
		return
	}*/
	GT = template.New("index").Funcs(tempower.FMap())
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
	http.HandleFunc("/newfamily", HandleNewFamily)
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/addmember", HandleAddMember)
	http.HandleFunc("/logout", HandleLogout)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
