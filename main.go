package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/coderconvoy/dbase"
	"github.com/coderconvoy/gojs"
	"github.com/coderconvoy/lazyf"
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

func CommonHandler() MuxFunc {
	cjs := []byte(CommonJS().Inner)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(cjs)
	}
}

func main() {
	insecure := flag.Bool("i", false, "Run without https")
	debug := flag.Bool("d", false, "Debug, log to fmt.Println")
	confloc := flag.String("c", "", "Configuration File Location")
	noconf := flag.Bool("noconf", false, "Run with no config file")
	flag.Parse()
	if *debug {
		dbase.SetQLogger(dbase.FmtLog{})
	}

	conf := lazyf.LZ{}
	if !*noconf {
		var err error
		conf, err = getConfig(*confloc)
		if err != nil {
			fmt.Println("Could not load config", err)
			return
		}
	}

	//Enable web access to embedded assets in this package
	gojs.Single.AddFuncs(Asset, AssetDir)

	//views
	http.HandleFunc("/", Handle)
	http.HandleFunc("/s/", HandleStatic)
	http.HandleFunc("/common.js", CommonHandler())
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
	//http.HandleFunc("/transactions", LoggedInVTemp(PageTransactions))
	http.HandleFunc("/personal", LoggedInView(PagePersonal))
	http.HandleFunc("/family", LoggedInView(PageFamily))
	//http.HandleFunc("/view", LoggedInView(Pag))

	dbase.QLog("Starting Server")

	if *insecure {
		log.Fatal(http.ListenAndServe(":8080", nil))
		return
	}

	pubkey := conf.PStringD("data/server.pub")
	privkey := conf.PStringD("data/server.key")

	log.Fatal(http.ListenAndServeTLS(":8081", pubkey, privkey, nil))
}
