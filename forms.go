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

func FormAddAccount(ld LoginData) *htmk.Tag {
	return htmk.NewParent("form", []*htmk.Tag{
		htmk.NewTag("input", "type", "text", "name", "username", "value", ld.Fmem, "--readonly", "--hidden"),
		htmk.NewTag("input", "type", "text", "name", "accountname", "pattern", ".{4,20}", "--required"),
		HTMLColorPicker("Col1", []string{}),
		HTMLColorPicker("Col2", []string{}),
		htmk.NewTextTag("div", "black:black", "class", "pocket", "id", "color_pocket"),
		htmk.NewTag("input", "type", "submit", "value", "add"),
	}, "id", "frm_add_account", "action", "addaccount", "method", "post")
}
