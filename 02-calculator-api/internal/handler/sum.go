package handler

import (
	"net/http"
)

type sumRequest []int

func (sr sumRequest) checkValues() error {
	return nil
}

func (sr sumRequest) execute() (result int) {
	for _, val := range sr {
		result += val
	}
	return result
}

func (h handler) Sum(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, &sumRequest{})
}
