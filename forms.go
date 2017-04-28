//This file is intended as forms, for creating html, replacing the former template system
package main

import (
	"time"

	htmk "github.com/coderconvoy/htmlmaker"
)

func HTMLColorPicker(id string, colors []string) *htmk.Tag {
	if len(colors) == 0 {
		colors = []string{"black", "blue", "red", "green", "orange", "white", "yellow"}
	}
	ops := []*htmk.Tag{}
	for _, v := range colors {
		ops = append(ops, htmk.NewTextTag("option", v, "value", v))
	}

	return htmk.NewParent("select", ops, "name", id, "id", id)
}

func SelectMyAccounts(ld LoginData, tagname string) *htmk.Tag {
	ops := []*htmk.Tag{}
	for _, v := range ld.Fam.ListWriteAccess(ld.Fmem) {
		s := v.Username + ":" + v.Name
		ops = append(ops, htmk.NewTextTag("option", s+": "+PrintMoney(v.End), "value", s))
	}
	return htmk.NewParent("select", ops, "name", tagname)
}

func SelectAllAccounts(ld LoginData, tagname string) *htmk.Tag {
	ops := []*htmk.Tag{}
	for _, v := range ld.Fam.Period.Accounts {
		s := v.Username + ":" + v.Name
		ops = append(ops, htmk.NewTextTag("option", s, "value", s))
	}
	return htmk.NewParent("select", ops, "name", tagname)
}

func SelectAllUsers(ld LoginData, tagname string) *htmk.Tag {
	ops := []*htmk.Tag{}
	for _, v := range ld.Fam.Members {
		ops = append(ops, htmk.NewTextTag("option", v.Username, "value", v.Username))
	}
	return htmk.NewParent("select", ops, "name", tagname)
}

func FormAddAccount(ld LoginData) *htmk.Tag {
	return htmk.NewParent("form", []*htmk.Tag{
		htmk.NewTextTag("h3", "Add Account"),
		htmk.QInput("text", "username", "value", ld.Fmem, "--readonly", "--hidden"),
		htmk.QInput("text", "accountname", "pattern", ".{4,20}", "--required"),
		HTMLColorPicker("Col1", []string{}),
		HTMLColorPicker("Col2", []string{}),
		htmk.NewTextTag("div", "black:black", "class", "pocket", "id", "color_pocket"),
		htmk.QSubmit("Add"),
	}, "id", "frm_add_account", "action", "addaccount", "method", "post")
}

func FormPay(ld LoginData) *htmk.Tag {
	return htmk.NewParent("form", []*htmk.Tag{
		htmk.NewTextTag("h3", "Pay Someone"),
		htmk.QInput("text", "username", "value", ld.Fmem, "--readonly", "--hidden"),
		htmk.NewText("<br>From:"), SelectMyAccounts(ld, "from"),
		htmk.NewText("<br>To:"), SelectAllAccounts(ld, "to"),
		htmk.NewText("<br>Amount:"), htmk.QInput("number", "amount", "step", "0.01", "min", "0"),
		htmk.NewText("<br>Purpose:"), html.QInput("text", "purpose"),
		html.QSubmit("Pay Now"),
	}, "id", "frm_pay", "action", "pay", "method", "post")
}

func FormStanding(ld LoginData) *htmk.Tag {
	return htmk.NewParent("form", []*htmk.Tag{
		htmk.NewTextTag("h3", "Make a Standing Order"),
		htmk.NewText("<br>From:"), SelectMyAccounts(ld, "from"),
		htmk.NewText("<br>To:"), SelectAllAccounts(ld, "to"),
		htmk.NewText("<br>Amount: £"), htmk.NewTag("input", "type", "number", "step", "0.01", "min", "0", "name", "amount"),
		htmk.NewText("<br>Purpose:"), htmk.NewTag("type", "text", "name", "purpose"),
		htmk.NewText("<br>Start Date:"), htmk.NewTag("type", "date", "name", "start", "value", time.Now().Format("2006-01-02")),
		htmk.NewText("<br>Then Every:"), htmk.NewTag("type", "number", "min", "1", "name", "delay", "value", "7"),
		htmk.QSelect("delay_type", "days", "months"),
		htmk.NewText("<br>"), htmk.QSubmit("Create"),
	}, "id", "frm_standing", "action", "addstanding", "method", "post")

}
