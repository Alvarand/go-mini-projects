package handler

import (
	"errors"
	"net/http"
)

var errorBadDivisor = errors.New("divisor is zero")

type divideRequest struct {
	Dividend *int `json:"dividend"`
	Divisor  *int `json:"divisor"`
}

func (dr divideRequest) checkValues() error {
	if dr.Dividend == nil || dr.Divisor == nil {
		return errorMissingValues
	}
	if *dr.Divisor == 0 {
		return errorBadDivisor
	}
	return nil
}

func (dr divideRequest) execute() int {
	return *dr.Dividend / *dr.Divisor
}

func (h handler) Divide(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, &divideRequest{})
}
