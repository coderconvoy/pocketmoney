package main

import (
	"errors"
	"path"
	"reflect"
	"text/template"
	"time"

	"github.com/coderconvoy/dbase"
	"github.com/coderconvoy/gojs"
	"github.com/coderconvoy/pocketmoney/history"
)

func Plex(p, a, b interface{}) interface{} {
	if p == nil || p == 0 || p == "" {
		return b
	}
	return a
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

func LoadAsset(f string) (string, error) {
	p := path.Join("assets", f)
	r, err := gojs.Single.Asset(p)
	return string(r), err
}

func SafeAsset(f string) string {
	p := path.Join("assets", f)
	r, err := gojs.Single.Asset(p)
	if err != nil {
		dbase.QLog("No Asset : " + p)
		return "//Asset not Found : " + p
	}
	return string(r)
}

func LoadJSAsset(f string) (string, error) {
	p := path.Join("assets/js", f)
	b, err := gojs.Single.Asset(p)
	return "<script>" + string(b) + "</script>", err
}

func JSEscape(f interface{}) (string, error) {
	switch v := f.(type) {
	case string:
		return template.JSEscapeString(v), nil
	case []byte:
		return template.JSEscapeString(string(v)), nil
	}
	return "", errors.New("Expected string or []byte")
}
