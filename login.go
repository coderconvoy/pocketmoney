package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/coderconvoy/dbase2"
)

var loginControl = NewLoginControl(time.Minute * 20)

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
func LoggedInFunc(f LoggedFunc, lock bool) MuxFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Fam, fmem, err := LoggedInFamily(w, r)
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
	}
}

// Logged In Family returns the loaded family file the family in the cookie id.
// TODO add boolean getLock
func LoggedInFamily(w http.ResponseWriter, r *http.Request) (*Family, string, error) {
	ld, iok := loginControl.GetLogin(w, r)
	if iok != dbase2.OK {
		return nil, "", errors.New("No login")
	}
	fam, err := LoadFamily(ld.Familyname)
	return fam, ld.Username, err
}
