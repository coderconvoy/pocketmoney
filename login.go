package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/coderconvoy/dbase2"
)

var loginControl = NewLoginControl(time.Minute * 20)
var logLock = dbase2.NewLocker()

type LoginStore struct {
	Familyname string
	Username   string
}

type LoginControl struct {
	*dbase2.SessionControl
}

func NewLoginControl(md time.Duration) *LoginControl {
	return &LoginControl{
		SessionControl: dbase2.NewSessionControl(md),
	}
}

func (lc *LoginControl) Login(w http.ResponseWriter, familyname, username string) {
	lc.SessionControl.Login(w, LoginStore{familyname, username})
}

func (lc *LoginControl) GetLogin(w http.ResponseWriter, r *http.Request) (LoginStore, int) {
	a, rtype := lc.SessionControl.GetLogin(w, r)
	if rtype != dbase2.OK {
		return LoginStore{}, rtype
	}
	return a.Data.(LoginStore), rtype
}

type LoggedFunc func(ld LoginData)

type MuxFunc func(w http.ResponseWriter, r *http.Request)

//Logged
func LoggedInFunc(f LoggedFunc, edit bool) MuxFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Fam, fmem, lockId, err := loggedInFamily(w, r)
		if !edit {
			logLock.Unlock(lockId)
		}
		if err != nil {
			GoIndex(w, r, err.Error())
			return
		}
		ld := LoginData{
			W:    w,
			R:    r,
			Fam:  Fam,
			Fmem: fmem,
		}
		f(ld)
		if edit {
			err := Fam.Save()
			if err != nil {
				fmt.Println("Save Error:", err)
			}
			logLock.Unlock(lockId)
		}

	}
}

// Logged In Family returns the loaded family file the family in the cookie id.
// TODO add boolean getLock
func loggedInFamily(w http.ResponseWriter, r *http.Request) (*Family, string, uint64, error) {
	ld, iok := loginControl.GetLogin(w, r)
	if iok != dbase2.OK {
		return nil, "", 0, errors.New("No login")
	}
	id := logLock.Lock(ld.Familyname)
	fam, err := LoadFamily(ld.Familyname)
	if err != nil {
		logLock.Unlock(id)
	}

	return fam, ld.Username, id, err
}
