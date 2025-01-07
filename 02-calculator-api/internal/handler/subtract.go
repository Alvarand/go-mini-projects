package handler

import (
	"net/http"
)

type subractRequest struct {
	baseRequest
}

func (sr subractRequest) execute() int {
	return *sr.Number1 - *sr.Number2
}

func (h handler) Subtract(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, &subractRequest{})
}
