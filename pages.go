package main

import (
	"github.com/coderconvoy/dbase"
	"github.com/coderconvoy/gojs"
	"github.com/coderconvoy/htmq"
)

func PageBasic(ld PageData, title string) (*htmq.Tag, *htmq.Tag) {
	fam, fmem := ld.Fam, ld.Fmem

	p, body := htmq.NewPage(title, "/s/main.css", "/common.js,https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js,")
	body.SetAttr("id", "main-area")

	banner := htmq.NewTag("div", "class", "banner")
	tmenu := htmq.NewTag("div", "class", "menu")

	//Add content to menu bar based on login status
	if fam != nil {
		tmenu.AddChildren(htmq.QLinkRep("/personal", "Personal"))
		if fam.IsParent(fmem) {
			tmenu.AddChildren(
				htmq.QLinkRep("/family", "Family"),
				htmq.QLinkRep("/transactions", "Transaction History"),
			)
		}
		tmenu.AddChildren(htmq.QLinkRep("/logout", "Logout"))
	}
	banner.AddChildren(
		htmq.QImg("/s/svg/banner.svg"),
		tmenu,
	)

	body.AddChildren(
		banner,
		htmq.NewTag("div", "style", "clear:both;"),
	)

	//hml := p.GetFirst(htmq.ByType("html"), 10)
	//if hml != nil {
	//}
	return p, body

}

func JSCalls() *htmq.Tag {
	return htmq.QScript(
		`
	showform();
	divstopocket(psvg);
	TallFrac();
		`)
	/*$("form").on("submit",function(event){
		event.preventDefault();
		this.submit();
		console.log(this);
	});*/
	/* $( "form" ).on( "submit", function( event ) {
	  	event.preventDefault();

		$.ajax({
			url:$(this).attr("action"),
			type:"post",
			data:$(this).serialize(),
			success:function(data){
				console.log("Success:" ,data);
			},
		});

	    console.log( $( this ).serialize() );
	*/
}

func CommonJS() *htmq.Tag {
	res, err := htmq.AScript(
		gojs.Single,
		"assets/js/showform.js",
		"assets/js/template.js",
		"assets/js/divtopocket.js",
		"assets/js/tallwide.js",
		"assets/js/login.js",
		`--psvg = `+"`"+SafeAsset("s/svg/pocket-temp.svg")+"`;",
	)
	if err != nil {
		dbase.QLog(err.Error())
	}
	return res

	/*		`
	showform("view_members");
	psvg = "`+SafeAsset("s/svg/pocket-temp.svg")+`;
	divstopocket(psvg);`,
		)*/
}

func PageFamily(ld PageData) *htmq.Tag {
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
			JSCalls(),
		}, "id", "allforms"),
	)
	return p
}

func PagePersonal(ld PageData) *htmq.Tag {

	//Side Buttons
	fbuts := htmq.NewParent("div", []*htmq.Tag{
		htmq.QBut("View Accounts", `showform('view_accounts')`),
		htmq.QBut("Add Account", `showform('frm_add_account')`),
		htmq.QBut("", `showform('frm_pay')`, "!/s/svg/payments.svg", "^Pay Someone"),
		htmq.QBut("", `showform('frm_request')`, "!/s/svg/requests.svg", "^Request Money"),
		htmq.QBut("Setup Regular Payment", `showform('frm_standing')`),
		htmq.QBut("Change Password", `showform('frm_pass')`),
		htmq.QBut("View Transactions", `showform('view_trans')`),
	}, "id", "actionlist")
	//Get Forms
	fl := htmq.NewParent("div", []*htmq.Tag{
		ViewAccounts(ld, "view_accounts"),
		ViewTransactions(ld, "view_trans"),
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
			JSCalls(),
		}, "id", "allforms"),
	)
	return p
}
