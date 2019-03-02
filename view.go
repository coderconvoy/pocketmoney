package main

import (
	"github.com/coderconvoy/htmq"
	"github.com/coderconvoy/money"
	"github.com/coderconvoy/pocketmoney/history"
)

func ViewTransactions(ld PageData, fid string) *htmq.Tag {

	rows := []*htmq.Tag{htmq.QMulti("tr", "th", "From", "To", "Amount", "Date", "Purpose")}
	for _, v := range ld.Fam.Period.Transactions {
		rows = append(rows, htmq.QMulti("tr", "td", v.From.String(), v.Dest.String(), v.Amount.String(), v.Date.Format("2006-01-02"), v.Purpose))
	}
	return htmq.NewParent("div", []*htmq.Tag{
		htmq.NewTextTag("h3", "Transactions"),
		htmq.NewParent("table", rows),
	}, "id", fid)

}

func ViewMyTransactions(ld PageData, fid string) *htmq.Tag {
	totals := make(map[history.ACKey]money.M)

	uname := ld.Fmem
	rows := []*htmq.Tag{htmq.QMulti("tr", "th", "From", "To", "Amount", "Total", "Date", "Purpose")}
	for _, v := range ld.Fam.Period.Transactions {

		if v.From.Username == uname {
			t, _ := totals[v.From]
			totals[v.From] = t - v.Amount
			rows = append(rows, htmq.QMulti("tr", "td", v.From.String(), v.Dest.String(), v.Amount.String(), totals[v.From].String(), v.Date.Format("2006-01-02"), v.Purpose))
		}

		if v.Dest.Username == uname {
			t, _ := totals[v.Dest]
			totals[v.Dest] = t + v.Amount
			rows = append(rows, htmq.QMulti("tr", "td", v.From.String(), v.Dest.String(), v.Amount.String(), totals[v.Dest].String(), v.Date.Format("2006-01-02"), v.Purpose))
		} else {
			continue
		}

	}
	return htmq.NewParent("div", []*htmq.Tag{
		htmq.NewTextTag("h3", "Transactions"),
		htmq.NewParent("table", rows),
	}, "id", fid)
}

func ViewMembers(ld PageData, fid string) *htmq.Tag {

	//List of li one for each family member
	rows := []*htmq.Tag{}
	for _, v := range ld.Fam.Members {
		rows = append(rows, htmq.NewParent("li", []*htmq.Tag{
			htmq.NewTextTag("p", v.Username+"  "+(Plex(v.Parent, "Parent", "Child")).(string)),
			viewMemberAccount(ld, v.Username),
		}))
	}

	return htmq.NewParent("ul", rows).Wrap("div", "id", fid)
}

func ViewAccounts(ld PageData, fid string) *htmq.Tag {
	return htmq.NewParent("div", []*htmq.Tag{
		viewMemberAccount(ld, ld.Fmem),
	}, "id", fid)
}

//ViewMemberAccount returns a ul containing all accounts for that user
func viewMemberAccount(ld PageData, uname string) *htmq.Tag {
	res := []*htmq.Tag{}
	for _, v := range ld.Fam.Period.UserAccounts(uname) {
		res = append(res, htmq.NewParent("li", []*htmq.Tag{
			htmq.NewTextTag("div", v.Col1+":"+v.Col2, "class", "pocket"),
			htmq.NewTextTag("div", v.Name+" "+v.End.String(), "class", "pocket-d"),
		}, "class", "pocket-li"))
	}
	return htmq.NewParent("ul", res)
}
