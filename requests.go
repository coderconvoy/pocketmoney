package main

import (
	"strconv"
	"time"
)

func HandleMakeRequest(ld *PageHand) (string, string) {
	fam := ld.Fam

	bt, err := readPostTransaction(ld)
	if err != nil {
		return "/personal", err.Error()
	}
	bt.Date = time.Now()

	fam.Requests = append(fam.Requests,
		&PaymentRequest{
			Transaction: bt,
			ID:          fam.NewRequestID(),
			Requester:   ld.Fmem,
			Returns:     0,
		})

	return "/personal", "Request Added"
}

func HandleRespondRequest(ld *PageHand) (string, string) {
	fam, r := ld.Fam, ld.R

	act := r.FormValue("action")
	rmid64, err := strconv.ParseInt(ld.R.FormValue("id"), 10, 32)
	if err != nil {
		return "/personal", "No ID given"
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
		return "/personal", "No Request exists with that ID"
	}

	switch act {
	case "accept":
		if req.From.Username != ld.Fmem {
			return "/personal", "Cannot accept someone else's request"
		}
		//Set date, and then add to transactions

		req.Date = time.Now()
		fam.Period.ApplyTransaction(req.Transaction)
		fam.Requests = append(fam.Requests[:remloc], fam.Requests[remloc+1:]...)

	case "reject":
		req.Returns++
	case "cancel":
		if req.Requester != ld.Fmem {
			return "/personal", "only requester can cancel a request"
		}
		fam.Requests = append(fam.Requests[:remloc], fam.Requests[remloc+1:]...)
	case "insist":
		req.Returns++
	}

	return "/personal", ""
}
