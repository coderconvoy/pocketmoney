package main

import (
	"strconv"
	"time"
)

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
	fam, w, r := ld.Fam, ld.W, ld.R

	act := r.FormValue("action")
	rmid64, err := strconv.ParseInt(ld.R.FormValue("id"), 10, 32)
	if err != nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd("No id Given"))
		return
	}
	id := int32(rmid64)
	var req *PaymentRequest = nil
	remloc := -1
	for i, rq := range ld.Fam.Requests {
		if rq.ID == id {
			req = rq
			remloc = i
			break
		}
	}
	if req == nil {
		ExTemplate(GT, w, "userhome.html", ld.Pd("No request of that ID"))
		return
	}

	switch act {
	case "accept":
		if req.From.Username != ld.Fmem {
			ExTemplate(GT, w, "userhome.html", ld.Pd("Cannot accept someone else's request"))
			return
		}
		//Set date, and then add to transactions

		req.Date = time.Now()
		fam.Period.ApplyTransaction(req.Transaction)
		fam.Requests = append(fam.Requests[:remloc], fam.Requests[remloc+1:]...)

	case "reject":
		req.Returns++
	case "cancel":
	case "insist":
	}

	ExTemplate(GT, w, "userhome.html", ld.Pd(""))
}
