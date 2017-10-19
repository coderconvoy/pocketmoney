package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/coderconvoy/dbase"
)

var loginControl = NewLoginControl(time.Minute * 20)
var logLock = dbase.NewLocker()

type LoginControl struct {
	*dbase.SessionControl
}

func NewLoginControl(md time.Duration) *LoginControl {
	return &LoginControl{
		SessionControl: dbase.NewSessionControl(md),
	}
}

func (lc *LoginControl) Login(w http.ResponseWriter, familyname, username string) LoginStore {
	ls := LoginStore{familyname, username, ""}
	lc.SessionControl.Login(w, ls)
	return ls
}

func (lc *LoginControl) GetLogin(w http.ResponseWriter, r *http.Request) (LoginStore, int) {
	a, rtype := lc.SessionControl.GetLogin(w, r)
	if rtype != dbase.OK {
		return LoginStore{}, rtype
	}
	return a.Data.(LoginStore), rtype
}

func (lc *LoginControl) EditLogin(r *http.Request, ls LoginStore) error {
	return lc.SessionControl.EditLogin(r, ls)
}

// PostFunc Performs the post operation on the given data
// Returns redirectloc , message
type PostFunc func(*PageHand) (string, string)

// ViewFunc Shows what the world looks like returning, the expected template name.
type DataFunc func(PageData, *http.Request) ([]byte, error)

type ViewFunc func(PageData) ([]byte, error)

type MuxFunc func(w http.ResponseWriter, r *http.Request)

func LoggedInData(f DataFunc) MuxFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pdata, lockId, err := loggedInFamily(w, r)
		if err != nil {
			GoIndex(w, r, err.Error())
			return
		}
		dbase.QLog(fmt.Sprintln("PData : ", pdata))
		//phand := &PageHand{PageData: pdata, W: w, R: r}
		if pdata.Fam.Calculate() {
			pdata.Fam.Save()
		}
		//Consider adding a calculate and save if changed here
		page, err := f(*pdata, r)
		if err != nil {
			http.Error(w, err.Error(), 400)
			dbase.QLog("Could not get data:" + err.Error())
		}

		_, err = w.Write(page)
		if err != nil {
			dbase.QLog("Could not write output to request")
		}

		pdata.Mes = ""
		err = loginControl.EditLogin(r, pdata.LoginStore)
		if err != nil {
			dbase.QLog("Could not edit login ")
		}
		logLock.Unlock(lockId)
	}

}

func LoggedInView(f ViewFunc) MuxFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pdata, lockId, err := loggedInFamily(w, r)
		if err != nil {
			GoIndex(w, r, err.Error())
			return
		}
		dbase.QLog(fmt.Sprintln("PData : ", pdata))
		//phand := &PageHand{PageData: pdata, W: w, R: r}
		if pdata.Fam.Calculate() {
			pdata.Fam.Save()
		}
		//Consider adding a calculate and save if changed here
		page, err := f(*pdata)
		if err != nil {
			http.Error(w, err.Error(), 400)
			dbase.QLog("Could not get data:" + err.Error())
		}

		_, err = w.Write(page)
		if err != nil {
			dbase.QLog("Could not write output to request")
		}

		pdata.Mes = ""
		err = loginControl.EditLogin(r, pdata.LoginStore)
		if err != nil {
			dbase.QLog("Could not edit login ")
		}
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
		phand := &PageHand{PageData: pdata, W: w, R: r}

		pdata.Fam.Calculate()

		path, mes := f(phand)
		pdata.Mes = mes

		err = pdata.Fam.Save()
		if err != nil {
			dbase.QLog("Save Error:" + err.Error())
		}

		dbase.QLog(fmt.Sprintln("Storing: ", pdata.LoginStore))
		err = loginControl.EditLogin(r, pdata.LoginStore)
		if err != nil {
			dbase.QLog("Could not edit login ")
		}

		logLock.Unlock(lockId)

		http.Redirect(w, r, path, 303)
	}
}

// Logged In Family returns the loaded family file the family in the cookie id.
func loggedInFamily(w http.ResponseWriter, r *http.Request) (*PageData, uint64, error) {
	ld, iok := loginControl.GetLogin(w, r)
	if iok != dbase.OK {
		return nil, 0, errors.New("No login")
	}
	id := logLock.Lock(ld.FamName)
	fam, err := LoadFamily(ld.FamName)
	return &PageData{fam, ld}, id, err
}
