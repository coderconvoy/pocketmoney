package main

import (
	"flag"
	"fmt"
	"net/http"
	"path"
	"text/template"

	"github.com/coderconvoy/dbase"
	"github.com/coderconvoy/gojs"
	"github.com/coderconvoy/templater/tempower"
)

var GT *template.Template
var FamDB = dbase.DBase{"data/families"}

type IndexData struct {
	Mes  string
	Logs []LoginPart
}

func (id IndexData) FamOptions() []Link {
	return []Link{}
}

func GoIndex(w http.ResponseWriter, r *http.Request, m string) {
	c, err := r.Cookie("LastLog")
	if err != nil {
		ExTemplate(GT, w, "index.html", IndexData{m, []LoginPart{}})
		return
	}
	var ll []LoginPart
	err = CookieUnmarshal(c.Value, &ll)
	if err != nil {
		ll = []LoginPart{}
	}

	ExTemplate(GT, w, "index.html", IndexData{m, ll})

}

func ExTemplate(t *template.Template, w http.ResponseWriter, name string, data interface{}) {
	err := t.ExecuteTemplate(w, name, data)
	if err != nil {
		fmt.Println(err)
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	dbase.QLog(r.Host + "--" + r.URL.Path)
	GoIndex(w, r, "")
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	loginControl.Logout(w, r)
	GoIndex(w, r, "You are now Logged out")
}

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	ass, err := Asset(path.Join("assets", p))
	if err != nil {
		dbase.QLog("Could not serve static, " + p + ":" + err.Error())
		return
	}
	switch path.Ext(p) {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "text/javascript")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	}
	w.Write(ass)
}

func main() {
	insecure := flag.Bool("i", false, "Run without https")
	debug := flag.Bool("d", false, "Debug, log to fmt.Println")
	flag.Parse()
	if *debug {
		dbase.SetQLogger(dbase.FmtLog{})
	}

	gojs.Single.AddFuncs(Asset, AssetDir)

	GT = template.New("index").Funcs(tempower.FMap()).Funcs(TemplateFuncs())
	ad, err := AssetDir("assets/templates")
	for _, n := range ad {
		if path.Ext(n) == ".swp" {
			continue
		}
		t, err := Asset("assets/templates/" + n)
		dbase.QLog("Parsing :" + n)
		GT = GT.New(n)
		_, err = GT.Parse(string(t))
		if err != nil {
			dbase.QLog(err.Error())
			fmt.Println(err)
			return
		}
	}

	http.HandleFunc("/", Handle)
	http.HandleFunc("/s/", HandleStatic)
	http.HandleFunc("/newfamily", HandleNewFamily)
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/logout", HandleLogout)

	//Edits
	http.HandleFunc("/addmember", LoggedInPost(HandleAddMember))
	http.HandleFunc("/addaccount", LoggedInPost(HandleAddAccount))
	http.HandleFunc("/pay", LoggedInPost(HandlePay))
	http.HandleFunc("/cancelstanding", LoggedInPost(HandleCancelStanding))
	http.HandleFunc("/makerequest", LoggedInPost(HandleMakeRequest))
	http.HandleFunc("/respondrequest", LoggedInPost(HandleRespondRequest))
	http.HandleFunc("/addstanding", LoggedInPost(HandleAddStanding))
	http.HandleFunc("/chpass", LoggedInPost(HandlePasswordChange))

	//Views
	http.HandleFunc("/transactions", LoggedInVTemp("transactions.html"))
	http.HandleFunc("/personal", LoggedInVTemp("userhome.html"))
	http.HandleFunc("/family", LoggedInVTemp("familypage.html"))
	http.HandleFunc("/view", LoggedInView(HandleViewAccount))

	if *insecure {
		err = http.ListenAndServe(":8080", nil)
		return

	}
	err = http.ListenAndServeTLS(":8081", "data/server.pub", "data/server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
}
