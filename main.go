package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	"time"

	dbase "github.com/coderconvoy/dbase2"
	"github.com/coderconvoy/templater/tempower"
)

var GT *template.Template
var FamDB = dbase.DBase{"data/families"}

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
	http.SetCookie(w, &http.Cookie{
		Name:    "Family",
		Value:   "DeleteCookie",
		Expires: time.Now().Add(-time.Second),
	})
	GT.ExecuteTemplate(w, "index.html", "You are now Logged out")

}

func HandleLoginFamily(w http.ResponseWriter, r *http.Request) {
	e := r.FormValue("Email")
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
	http.SetCookie(w, &http.Cookie{
		Name:  "Family",
		Value: f.Email,
	})
	ExTemplate(GT, w, "familypage.html", f)

}
func HandleNewFamily(w http.ResponseWriter, r *http.Request) {
	var f Family
	f.Email = r.FormValue("Email")
	f.FamilyName = r.FormValue("FamilyName")

	_, err := FamDB.ReadMap("Email")
	if err == nil {
		GT.ExecuteTemplate(w, "index.html", "Email Already Exists")
		return
	}

	mar, err := json.Marshal(&f)
	if err != nil {
		fmt.Println("could not marshal f:", err)
	}
	FamDB.WriteMap(f.Email, mar)
	http.SetCookie(w, &http.Cookie{
		Name:  "Family",
		Value: f.Email,
	})
	GT.ExecuteTemplate(w, "familypage.html", f)

}

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
}

func HandleNewAccount(w http.ResponseWriter, r *http.Request) {
	/*c, err := r.Cookie("Username")
	if err != nil {
		GT.ExecuteTemplate(w, "index.html", nil)
	}
	b, err := UserDB.ReadMap(c.Value)
	if err != nil {
		fmt.Fprintf(w, "No thing stored for user %s", c.Value)
	}

	GT.ExecuteTemplate(w, "newaccount.html", b)
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
	http.HandleFunc("/loginfamily", HandleLoginFamily)
	http.HandleFunc("/logout", HandleLogout)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
