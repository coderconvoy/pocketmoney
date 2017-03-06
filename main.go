package main

import (
	"flag"
	"fmt"
	"net/http"
	"path"
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

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	ass, err := Asset(path.Join("assets", p))
	if err != nil {
		fmt.Println("Could not serve static, ", p, ":", err)
		return
	}
	switch path.Ext(p) {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":

	}
	w.Write(ass)
}

func main() {
	insecure := flag.Bool("i", false, "Run without https")
	flag.Parse()

	GT = template.New("index").Funcs(tempower.FMap())
	ad, err := AssetDir("assets/templates")
	for _, n := range ad {
		if path.Ext(n) == ".swp" {
			continue
		}
		t, err := Asset("assets/templates/" + n)
		fmt.Println("Parsing :" + n)
		GT = GT.New(n)
		_, err = GT.Parse(string(t))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	http.HandleFunc("/s/", HandleStatic)
	http.HandleFunc("/newfamily", HandleNewFamily)
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/addmember", HandleAddMember)
	http.HandleFunc("/logout", HandleLogout)
	http.HandleFunc("/personal", HandlePersonal)
	http.HandleFunc("/addaccount", HandleAddAccount)
	http.HandleFunc("/transactions", HandleTransactions)
	http.HandleFunc("/family", HandleFamily)
	http.HandleFunc("/pay", HandlePay)
	http.HandleFunc("/chpass", HandlePasswordChange)
	http.HandleFunc("/", Handle)

	if *insecure {
		err = http.ListenAndServe(":8080", nil)
		return

	}
	err = http.ListenAndServeTLS(":8081", "data/server.pub", "data/server.key", nil)
	if err != nil {
		fmt.Println(err)
	}

}
