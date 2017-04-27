//This file is intended as forms, for creating html, replacing the former template system
package main

import (
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
		htmk.NewTag("input", "type", "text", "name", "username", "value", ld.Fmem, "--readonly", "--hidden"),
		htmk.NewTag("input", "type", "text", "name", "accountname", "pattern", ".{4,20}", "--required"),
		HTMLColorPicker("Col1", []string{}),
		HTMLColorPicker("Col2", []string{}),
		htmk.NewTextTag("div", "black:black", "class", "pocket", "id", "color_pocket"),
		htmk.NewTag("input", "type", "submit", "value", "add"),
	}, "id", "frm_add_account", "action", "addaccount", "method", "post")
}

func FormPay(ld LoginData) *htmk.Tag {
	return htmk.NewParent("form", []*htmk.Tag{
		htmk.NewTextTag("h3", "Pay Someone"),
		htmk.NewTag("input", "type", "text", "name", "username", "value", ld.Fmem, "--readonly", "--hidden"),
		htmk.NewText("<br>From:"), SelectMyAccounts(ld, "from"),
		htmk.NewText("<br>To:"), SelectAllAccounts(ld, "to"),
		htmk.NewText("<br>Amount:"), htmk.NewTag("input", "type", "number", "step", "0.01", "min", "0", "name", "amount"),
		htmk.NewText("<br>Purpose:"), html.NewTag("input", "type", "text", "name", "purpose"),
		html.NewTag("input", "type", "submit", "value", "Pay Now"),
	}, "id", "frm_pay", "action", "pay", "method", "post")
}
