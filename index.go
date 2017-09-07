package main

import (
	"fmt"
	"net/http"

	"github.com/coderconvoy/htmq"
)

func GoIndex(w http.ResponseWriter, r *http.Request, m string) {
	c, err := r.Cookie("LastLog")
	var ll []LoginStore

	if err != nil {
		w.Write(PageIndex(m, ll).Bytes())
		return
	}
	err = CookieUnmarshal(c.Value, &ll)
	if err != nil {
		ll = []LoginStore{}
	}

	w.Write(PageIndex(m, ll).Bytes())

}

func PageIndex(mes string, ll []LoginStore) *htmq.Tag {
	page, body := PageBasic(PageData{}, "Pocket Money")
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
			v.FamName+"<br>"+v.Fmem,
			fmt.Sprintf("showlogin('%s','%s')", v.FamName, v.Fmem),
		))
	}
	l0 := LoginStore{"", "", nil}
	if len(ll) > 0 {
		l0 = ll[0]
	}

	//Login Form
	lForm := htmq.QForm("login", []*htmq.Tag{
		htmq.NewTextTag("h2", "Login"),
		htmq.NewText("Family Name : "), htmq.QInput("text", "family", "id", "linfam", "value", l0.FamName),
		htmq.NewText("<br>User Name : "), htmq.QInput("text", "username", "id", "linusr", "value", l0.Fmem),
		htmq.NewText("<br>Password : "), htmq.QInput("password", "pwd", "id", "linpass"),
		htmq.QSubmit("Login"),
	}, "id", "login")

	//New Family Form
	nForm := htmq.QForm("newfamily", []*htmq.Tag{
		htmq.NewTextTag("h2", "Create New Family"),
		htmq.QText("Family Name : "), htmq.QInput("text", "familyname", "id", "famtop"),
		htmq.QText("<br>User Name : "), htmq.QInput("text", "username"),
		htmq.QText("<br>Email : "), htmq.QInput("email", "email"),
		htmq.NewText("<br>Password : "), htmq.QInput("password", "pwd1", "id", "linpass"),
		htmq.NewText("<br>Confirm : "), htmq.QInput("password", "pwd2", "id", "linpass"),
		htmq.QSubmit("Create Family"),
	}, "id", "newfamily")

	flist := htmq.NewTag("div", "id", "formlist")
	flist.AddChildren(lForm, nForm)
	body.AddChildren(aList, flist, JSCalls())
	return page
}
