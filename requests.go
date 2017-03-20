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
	w, r := ld.W, ld.R

	act := r.FormValue("action")
	id := r.FormValue("id")
	var req PaymentRequest = nil
	for _, rq := range ld.Fam.Requests {
		if rq.ID == id {
			req = rq
			break
		}
	}
	if req == nil {
		ExTemplate(GT, ld.W, "userhome.html", ld.PD("No request of that ID"))
		return
	}

	switch act {
	case "accept":
		if req.From.Username != ld.Fmem {
			ExTemplate(GT, ld.W, "userhome.html", ld.PD("Cannot accept someone else's request"))
			return
		}
		//Todo set date, and then add to transactions

	case "reject":
	case "cancel":
	case "insist":
	}

	ExTemplate(GT, ld.W, "userhome.html", ld.Pd("Response handler not ready"))
}
