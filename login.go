package main

import (
	"net/http"
	"time"

	"github.com/coderconvoy/dbase2"
)

var loginControl = NewLoginControl(time.Minute * 20)

type LoginData struct {
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
	lc.SessionControl.Login(w, LoginData{familyname, username})
}

func (lc *LoginControl) GetLogin(w http.ResponseWriter, r *http.Request) (LoginData, int) {
	a, rtype := lc.SessionControl.GetLogin(w, r)
	if rtype != dbase2.OK {
		return LoginData{}, rtype
	}
	return a.Data.(LoginData), rtype
}
