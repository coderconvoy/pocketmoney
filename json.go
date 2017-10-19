package main

import "net/http"

func JsonHistory(ld PageData, r *http.Request) ([]byte, error) {
	a := r.FormValue("POOP")
	return []byte("HELLOPOO," + a), nil
}
