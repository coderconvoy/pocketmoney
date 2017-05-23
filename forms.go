//This file is intended as forms, for creating html, replacing the former template system
package main

import (
	"time"

	"github.com/coderconvoy/htmq"
)

func HTMLColorPicker(id string, colors []string) *htmq.Tag {
	if len(colors) == 0 {
		colors = []string{"black", "blue", "red", "green", "orange", "white", "yellow"}
	}
	ops := []*htmq.Tag{}
	for _, v := range colors {
		ops = append(ops, htmq.NewTextTag("option", v, "value", v))
	}

	return htmq.NewParent("select", ops, "name", id, "id", id)
}

func SelectMyAccounts(ld PageData, tagname string) *htmq.Tag {
	ops := []*htmq.Tag{}
	for _, v := range ld.Fam.ListWriteAccess(ld.Fmem) {
		s := v.Username + ":" + v.Name
		ops = append(ops, htmq.NewTextTag("option", s+": "+v.End.String(), "value", s))
	}
	return htmq.NewParent("select", ops, "name", tagname)
}

func SelectAllAccounts(ld PageData, tagname string) *htmq.Tag {
	ops := []*htmq.Tag{}
	for _, v := range ld.Fam.Period.Accounts {
		s := v.Username + ":" + v.Name
		ops = append(ops, htmq.NewTextTag("option", s, "value", s))
	}
	return htmq.NewParent("select", ops, "name", tagname)
}

func SelectAllUsers(ld PageData, tagname string) *htmq.Tag {
	ops := []*htmq.Tag{}
	for _, v := range ld.Fam.Members {
		ops = append(ops, htmq.NewTextTag("option", v.Username, "value", v.Username))
	}
	return htmq.NewParent("select", ops, "name", tagname)
}

func FormAddAccount(ld PageData) *htmq.Tag {
	return htmq.NewParent("form", []*htmq.Tag{
		htmq.NewTextTag("h3", "Add Account"),
		htmq.QInput("text", "accountname", "pattern", ".{4,20}", "--required"),
		HTMLColorPicker("Col1", []string{}),
		HTMLColorPicker("Col2", []string{}),
		htmq.NewTextTag("div", "black:black", "class", "pocket", "id", "color_pocket"),
		htmq.QSubmit("Add"),
	}, "id", "frm_add_account", "action", "addaccount", "method", "post")
}

func FormPay(ld PageData) *htmq.Tag {
	return htmq.QForm("pay", []*htmq.Tag{
		htmq.NewTextTag("h3", "Pay Someone"),
		htmq.NewText("<br>From:"), SelectMyAccounts(ld, "from"),
		htmq.NewText("<br>To:"), SelectAllAccounts(ld, "to"),
		htmq.NewText("<br>Amount:"), htmq.QInput("number", "amount", "step", "0.01", "min", "0"),
		htmq.NewText("<br>Purpose:"), htmq.QInput("text", "purpose"),
		htmq.QSubmit("Pay Now"),
	}, "id", "frm_pay")
}

func FormStanding(ld PageData) *htmq.Tag {
	return htmq.QForm("addstanding", []*htmq.Tag{
		htmq.NewTextTag("h3", "Make a Standing Order"),
		htmq.NewText("<br>From:"), SelectMyAccounts(ld, "from"),
		htmq.NewText("<br>To:"), SelectAllAccounts(ld, "to"),
		htmq.NewText("<br>Amount: £"), htmq.QInput("number", "amount", "step", "0.01", "min", "0"),
		htmq.NewText("<br>Purpose:"), htmq.QInput("text", "purpose"),
		htmq.NewText("<br>Start Date:"), htmq.QInput("date", "start", "value", time.Now().Format("2006-01-02")),
		htmq.NewText("<br>Then Every:"), htmq.QInput("number", "delay", "min", "1", "value", "7"),
		htmq.QSelect("delay_type", "days", "months"),
		htmq.NewText("<br>"), htmq.QSubmit("Create"),
	}, "id", "frm_standing")

}

func FormPassword() *htmq.Tag {
	return htmq.QForm("chpass", []*htmq.Tag{
		htmq.NewTextTag("h3", "Change Password"),
		htmq.NewText("Old Password : "), htmq.QInput("password", "oldpwd"),
		htmq.NewText("<br>New Password : "), htmq.QInput("password", "pwd1"),
		htmq.NewText("<br>Confirm : "), htmq.QInput("password", "pwd2"),
		htmq.QSubmit("Change"),
	}, "id", "frm_chpass")
}

func FormRequest(ld PageData) *htmq.Tag {
	return htmq.QForm("makerequest", []*htmq.Tag{
		htmq.NewTextTag("h3", "Request A payment"),
		htmq.NewTextTag("p", "The owner of the sending account will have to authorise this payment"),
		htmq.NewText("From : "), SelectAllUsers(ld, "fromuser"),
		htmq.NewText("<br>To : "), SelectAllAccounts(ld, "to"),
		htmq.NewText("<br>Amount : "), htmq.QInput("number", "amount", "step", "0.01", "min", "0"),
		htmq.NewText("<br>Purpose : "), htmq.QInput("text", "purpose"),
		htmq.QSubmit("Request Now"),
	}, "id", "frm_request")
}

func FormAddMember() *htmq.Tag {
	return htmq.QForm("addmember", []*htmq.Tag{
		htmq.NewTextTag("h3", "Add a Family Member"),
		htmq.NewText("Name : "), htmq.QInput("text", "username"),
		htmq.NewText("<br>Is Parent : "), htmq.QInput("checkbox", "parent"),
		htmq.NewText("<br>Password : "), htmq.QInput("password", "pwd1", "pattern", ".{5,20}"),
		htmq.NewText("<br>(Must Contain at least 5 Characters)"),
		htmq.NewText("<br>Confirm : "), htmq.QInput("password", "pwd2", "pattern", ".{5,20}"),
		htmq.QSubmit("Add"),
	}, "id", "frm_add_member")
}
