package main

import "time"

func HandleMakeRequest(ld LoginData) {
	w, fam := ld.W, ld.Fam

	bt, err := readPostTransaction(ld)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd(err.Error()))
		return
	}
	bt.Date = time.Now()

	fam.Requests = append(fam.Requests,
		&PaymentRequest{
			Transaction: bt,
			ID:          fam.NewRequestID(),
			Requester:   ld.Fmem,
			Returns:     0,
		})

	ExTemplate(GT, w, "userhome.html", ld.Pd("Request Added"))
}

func HandleRespondRequest(ld LoginData) {
	ExTemplate(GT, ld.W, "userhome.html", ld.Pd("Response handler not ready"))
}
