package main

import "testing"

func Test_cMarshal(t *testing.T) {
	logs := []LoginStore{
		{"f1", "fm1"},
		{"f2", "fm1"},
		{"f3", "fm2"},
		{"f3", "fm1"},
	}
	s, err := CookieMarshal(logs)
	if err != nil {
		t.Log(err)
	}

	var res []LoginStore

	err = CookieUnmarshal(s, &res)
	if err != nil {
		t.Log(err)
	}

	if len(logs) != len(res) {
		t.Logf("Lengths not matching, %d,%d", len(logs), len(res))
		t.Fail()
	}
}
