package main

import (
	"github.com/coderconvoy/dbase"
	"github.com/coderconvoy/htmq"
)

func PageBasic(ld LoginData, title string) (*htmq.Tag, *htmq.Tag) {
	fam, fmem := ld.Fam, ld.Fmem

	p, body := htmq.NewPage(title, "/s/main.css")
	body.SetAttr("id", "main-area")

	banner := htmq.NewTag("div", "class", "banner")
	tmenu := htmq.NewTag("div", "class", "menu")

	//Add content to menu bar based on login status
	if fam != nil {
		tmenu.AddChildren(htmq.QLink("/personal", "Personal"))
		if fam.IsParent(fmem) {
			tmenu.AddChildren(
				htmq.QLink("/family", "Family"),
				htmq.QLink("/transactions", "Transaction History"),
			)
		}
		tmenu.AddChildren(htmq.QLink("/logout", "Logout"))
	}
	banner.AddChildren(
		htmq.QImg("/s/svg/banner.svg"),
		tmenu,
	)

	twide, err := LoadAsset("tallwide.js")
	if err != nil {
		dbase.QLog("No Asset tallwide.js")
	}
	body.AddChildren(
		banner,
		htmq.NewTag("div", "style", "clear:both;"),

		htmq.NewTextTag("script", twide),
	)

	return p, body
}

func CommonJS() *htmq.Tag {
	return htmq.QScript(
		SafeAsset("js/showform.js"),
		SafeAsset("js/template.js"),
		SafeAsset("js/divtopocket.js"),
		`
showform("view_members");
psvg = "`+SafeAsset("s/svg/pocket-temp.svg")+`;
divstopocket(psvg);`,
	)
}

func PageFamily(ld LoginData) *htmq.Tag {
	fam := ld.Fam

	//Side Buttons
	fbuts := htmq.NewParent("div", []*htmq.Tag{
		htmq.QBut("View Members", `showform('viewmembers')`),
		htmq.QBut("Add Member", `showform('frm_add_member')`),
	}, "id", "actionlist")
	//Get Forms
	fl := htmq.NewParent("div", []*htmq.Tag{
		ViewMembers(ld, "viewmembers"),
		FormAddMember(),
	}, "id", "formlist")

	//Get page
	p, body := PageBasic(ld, "Family")

	//Add forms to page
	body.AddChildren(
		htmq.NewTextTag("h2", "Family Page : "+fam.FamilyName),
		htmq.NewParent("div", []*htmq.Tag{
			fbuts,
			fl,
			CommonJS(),
		}, "id", "allforms"),
	)
	return p
}

func PagePersonal(ld LoginData) *htmq.Tag {

	//Side Buttons
	fbuts := htmq.NewParent("div", []*htmq.Tag{
		htmq.QBut("View Accounts", `showform('view_accounts')`),
		htmq.QBut("Add Account", `showform('frm_add_account')`),
		htmq.QBut("", `showform('frm_pay')`, "!/s/svg/payments.svg", "^Pay Someone"),
		htmq.QBut("", `showform('frm_request')`, "!/s/svg/requests.svg", "^Request Money"),
		htmq.QBut("Setup Regular Payment", `showform('frm_standing')`),
		htmq.QBut("Change Password", `showform('frm_standing')`),
	}, "id", "actionlist")
	//Get Forms
	fl := htmq.NewParent("div", []*htmq.Tag{
		ViewAccounts(ld, "view_acounts"),
		FormAddAccount(ld),
		FormPay(ld),
		FormRequest(ld),
		FormStanding(ld),
		FormPassword(),
	}, "id", "formlist")

	//Get page
	p, body := PageBasic(ld, "Personal")

	//Add forms to page
	body.AddChildren(
		htmq.NewTextTag("h2", "Personal Page : "+ld.Fmem),
		htmq.NewParent("div", []*htmq.Tag{
			fbuts,
			fl,
			CommonJS(),
		}, "id", "allforms"),
	)
	return p
}
