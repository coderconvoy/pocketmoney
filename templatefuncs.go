package main

import (
	"fmt"
	"path"
	"reflect"
	"text/template"
	"time"

	"github.com/coderconvoy/pocketmoney/history"
)

func TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"plex":         Plex,
		"loginpart0":   LoginPart0,
		"standingbyac": FilterStandingByAC,
		"money":        PrintMoney,
		"date":         PrintDate,
		"dateRFC":      PrintDateRFC,
		"type":         PrintType,
		"eq2":          Eq2,
		"js":           LoadJSAsset,
	}
}

func Plex(p, a, b interface{}) interface{} {
	if p == nil || p == 0 || p == "" {
		return b
	}
	return a

}

func PrintMoney(n int) string {
	if n < 0 {
		return "-£" + fmt.Sprintf("%.2f", float32(-n)/100)
	}
	return "£" + fmt.Sprintf("%.2f", float32(n)/100)
}

func PrintDate(t ...time.Time) string {
	if len(t) == 0 {
		return time.Now().Format("Mon 2/Jan/06")
	}
	return t[0].Format("Mon 2/Jan/06")
}

func PrintDateRFC(t ...time.Time) string {
	if len(t) == 0 {
		return time.Now().Format("2006-01-02")
	}
	return t[0].Format("2006-01-02")
}

func FilterStandingByAC(st []*StandingOrder, ac history.Account) []*StandingOrder {
	res := []*StandingOrder{}
	for _, v := range st {
		if (ac.ACKey == v.From) || (ac.ACKey == v.Dest) {
			res = append(res, v)
		}
	}
	return res
}

func PrintType(o interface{}) string {
	return reflect.TypeOf(o).String()
}

func Eq2(a, b interface{}) bool {
	return a == b
}

func LoginPart0(lar []LoginPart) LoginPart {
	if len(lar) > 0 {
		return lar[0]
	}
	return LoginPart{}
}

func LoadJSAsset(f string) (string, error) {
	p := path.Join("assets/js", f)
	b, err := Asset(p)
	return "<script>" + string(b) + "</script>", err
}
