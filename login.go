package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/coderconvoy/dbase2"
)

var loginControl = NewLoginControl(time.Minute * 20)
var logLock = dbase2.NewLocker()

type LoginControl struct {
	*dbase2.SessionControl
}

func NewLoginControl(md time.Duration) *LoginControl {
	return &LoginControl{
		SessionControl: dbase2.NewSessionControl(md),
	}
}

func (lc *LoginControl) Login(w http.ResponseWriter, familyname, username string) {
	lc.SessionControl.Login(w, LoginStore{familyname, username, []JPar{}})
}

func (lc *LoginControl) GetLogin(w http.ResponseWriter, r *http.Request) (LoginStore, int) {
	a, rtype := lc.SessionControl.GetLogin(w, r)
	if rtype != dbase2.OK {
		return LoginStore{}, rtype
	}
	return a.Data.(LoginStore), rtype
}

func (lc *LoginControl) EditLogin(r *http.Request, ls LoginStore) error {
	return lc.SessionControl.EditLogin(r, ls)
}

// PostFunc Performs the post operation on the given data
// Returns redirectloc , jobs
type PostFunc func(PageData) (string, []JPar)

// ViewFunc Shows what the world looks like returning, the expected template name.
type ViewFunc func(PageData) string

type MuxFunc func(w http.ResponseWriter, r *http.Request)

func LoggedInView(tname string) MuxFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pdata, lockId, err := loggedInFamily(w, r)
		//Consider adding a calculate and save if changed here
		ExTemplate(GT, w, tname, pdata)
		logLock.Unlock(lockId)

	}
}

//Logged
func LoggedInPost(f PostFunc) MuxFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pdata, lockId, err := loggedInFamily(w, r)
		if err != nil {
			GoIndex(w, r, err.Error())
			return
		}
		path, jobs := f(pdata)

		err = pdata.Fam.Save()
		if err != nil {
			dbase2.QLog("Save Error:" + err.Error())
		}

		pp := pdata.LoginStore
		pp.Jobs = jobs
		err = loginControl.EditLogin(r, pp)

		logLock.Unlock(lockId)

		http.Redirect(w, r, path, 303)
	}
}

// Logged In Family returns the loaded family file the family in the cookie id.
func loggedInFamily(w http.ResponseWriter, r *http.Request) (PageData, uint64, error) {
	ld, iok := loginControl.GetLogin(w, r)
	if iok != dbase2.OK {
		return PageData{}, 0, errors.New("No login")
	}
	id := logLock.Lock(ld.Familyname)
	fam, err := LoadFamily(ld.Familyname)
	if err != nil {
		logLock.Unlock(id)
	}

	return PageData{Fam: fam, LoginStore: ld}, id, err
}
