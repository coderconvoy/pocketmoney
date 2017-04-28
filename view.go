package main

import (
	htmk "github.com/coderconvoy/htmlmaker"
)

func ViewTransactions(ld LoginData) *htmk.Tag {

	rows := []*htmk.Tag{htmk.QMulti("tr", "th", "From", "To", "Amount", "Date", "Purpose")}
	for _, v := range ld.Fam.Period.Transactions {
		rows = append(rows, htmk.QMulti("tr", "td", v.From.String(), v.Dest.String(), PrintMoney(v.Amount), v.Date.Format("2006-01-02"), v.Purpose))
	}
	res := htmk.NewParent("div", []*htmk.Tag{
		htmk.NewTextTag("h3", "Transactions"),
		htmk.NewParent("table", rows),
	})

	return res.Wrap("div", "id", "formlist")

}
