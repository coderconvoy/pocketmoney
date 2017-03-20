package main

import "time"

func HandleMakeRequest(ld LoginData) {
	w, r, fam := ld.W, ld.R, ld.Fam

	bt, err := readPostTransaction(ld)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd(err.Error()))
		return
	}
	bt.Date = time.Now()

	//	for

}

func HandleRespondRequest(ld LoginData) {
}
