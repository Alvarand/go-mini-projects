package handler

import (
	"net/http"
)

type addRequest struct {
	baseRequest
}

func (ar addRequest) execute() int {
	return *ar.Number1 + *ar.Number2
}

func (h handler) Add(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, &addRequest{})
}
