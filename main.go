package main

import (
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
	_insec := lazyf.FlagBool("i", "insec", "Run local without https")
	_debug := lazyf.FlagBool("d", "debug", "Debug, log to fmt.Println")
	_dataloc := lazyf.FlagString("df", "{HOME}/data/pocketmoney", "datafolder", "Location of data folder")
	_pubkey := lazyf.FlagString("pbk", "data/server.pub", "pubkey", "Location of public Key")
	_privkey := lazyf.FlagString("prk", "data/server.key", "privkey", "Location of Private Key")
	_port := lazyf.FlagString("p", "8080", "port", "Port to serv")

	lazyf.FlagLoad("c", "{HOME}/.config/pocketmoney/init.lz")

	FamDB.Root = path.Join(lazyf.EnvReplace(*_dataloc), "families")

	if *_debug {
		dbase.SetQLogger(dbase.FmtLog{})
	}

	//Enable web access to embedded assets in this package
	gojs.Single.AddFuncs(Asset, AssetDir)

	//basic views
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

	//Data Views
	//http.HandleFunc("/transactions", LoggedInVTemp(PageTransactions))
	http.HandleFunc("/personal", LoggedInView(PagePersonal))
	http.HandleFunc("/family", LoggedInView(PageFamily))
	http.HandleFunc("/json/history", LoggedInData(JsonHistory))
	//http.HandleFunc("/view", LoggedInView(Pag))

	dbase.QLog("Starting Server")

	if *_insec {
		log.Fatal(http.ListenAndServe("localhost:"+*_port, nil))
		return
	}

	log.Fatal(http.ListenAndServeTLS(":"+*_port, *_pubkey, *_privkey, nil))
}
