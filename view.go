package main

import "github.com/coderconvoy/htmq"

func ViewTransactions(ld LoginData, fid string) *htmq.Tag {

	rows := []*htmq.Tag{htmq.QMulti("tr", "th", "From", "To", "Amount", "Date", "Purpose")}
	for _, v := range ld.Fam.Period.Transactions {
		rows = append(rows, htmq.QMulti("tr", "td", v.From.String(), v.Dest.String(), PrintMoney(v.Amount), v.Date.Format("2006-01-02"), v.Purpose))
	}
	return htmq.NewParent("div", []*htmq.Tag{
		htmq.NewTextTag("h3", "Transactions"),
		htmq.NewParent("table", rows),
	}, "id", fid)

}

func ViewMembers(ld LoginData, fid string) *htmq.Tag {
	//membertag returns a ul containing all accounts for that user
	membertag := func(uname string) *htmq.Tag {
		res := []*htmq.Tag{}
		for _, v := range ld.Fam.Period.UserAccounts(uname) {
			res = append(res, htmq.NewParent("li", []*htmq.Tag{
				htmq.NewTextTag("div", v.Col1+":"+v.Col2, "class", "pocket"),
				htmq.NewTextTag("div", v.Name+" "+PrintMoney(v.End), "class", "pocket-d"),
			}, "class", "pocket-li"))
		}
		return htmq.NewParent("ul", res)
	}

	//List of li one for each family member
	rows := []*htmq.Tag{}
	for _, v := range ld.Fam.Members {
		rows = append(rows, htmq.NewParent("li", []*htmq.Tag{
			htmq.NewTextTag("p", v.Username+"  "+(Plex(v.Parent, "Parent", "Child")).(string)),
			membertag(v.Username),
		}))
	}

	return htmq.NewParent("ul", rows).Wrap("div", "id", fid)
}
