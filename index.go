package main

import (
	"fmt"
	"net/http"

	"github.com/coderconvoy/htmq"
)

func GoIndex(w http.ResponseWriter, r *http.Request, m string) {
	c, err := r.Cookie("LastLog")
	if err != nil {
		ExTemplate(GT, w, "index.html", IndexData{m, []LoginStore{}})
		return
	}
	var ll []LoginStore
	err = CookieUnmarshal(c.Value, &ll)
	if err != nil {
		ll = []LoginStore{}
	}

	ExTemplate(GT, w, "index.html", IndexData{m, ll})

}

func PageIndex(mes string, ll []LoginStore) {
	p, body := htmq.NewPage(title, "/s/main.css")
	body.SetAttr("id", "main-area")
	if mes != "" {
		body.AddChildren(htmq.NewTextTag("b", mes))
	}

	//Action Buttons inc named Buttons for previous logins
	aList := htmq.NewParent("div", []*htmq.Tag{
		htmq.QBut("New Family", "shownewfam()"),
		htmq.QBut("Login", "showlogin()"),
	}, "id", "actionlist")
	for _, v := range ll {
		aList.AddChildren(htmq.QBut(
			v.FamName+"<br>"+v.FMem,
			fmt.Sprintf("showlogin('%s','%s')", v.FamName, v.Fmem),
		))
	}
	l0 := LoginStore{"", "", nil}
	if len(ll) > 0 {
		l0 = ll[0]
	}

	//Login Form
	lForm := htmq.QForm("login", []htmq.Tag{
		htmq.NewTextTag("h2", "Login"),
		htmq.NewText("Family Name : "), htmq.QInput("text", "family", "id", "linfam", "value", l0.FamName),
		htmq.NewText("<br>User Name : "), htmq.QInput("text", "username", "id", "linusr", "value", l0.Fmem),
		htmq.NewText("<br>Password : "), htmq.QInput("password", "pwd", "id", "linpass"),
		htmq.QSubmit("Login"),
	}, "id", "login")

	//Family Form

}
