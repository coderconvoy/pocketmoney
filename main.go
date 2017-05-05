package main

import (
	"flag"
	"fmt"
	"net/http"
	"path"

	"github.com/coderconvoy/dbase"
	"github.com/coderconvoy/gojs"
)

var FamDB = dbase.DBase{"data/families"}

type IndexData struct {
	Mes  string
	Logs []LoginStore
}

func (id IndexData) FamOptions() []Link {
	return []Link{}
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
	http.HandleFunc("/transactions", LoggedInVTemp(PageTransactions))
	http.HandleFunc("/personal", LoggedInVTemp(PagePerson))
	http.HandleFunc("/family", LoggedInView(PageFamily))
	//http.HandleFunc("/view", LoggedInView(Pag))

	if *insecure {
		err = http.ListenAndServe(":8080", nil)
		return

	}
	err = http.ListenAndServeTLS(":8081", "data/server.pub", "data/server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
}
