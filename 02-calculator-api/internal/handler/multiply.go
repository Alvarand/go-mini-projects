package handler

import (
	"net/http"
)

type multiplyRequest struct {
	baseRequest
}

func (mr multiplyRequest) execute() int {
	return *mr.Number1 * *mr.Number2
}

func (h handler) Multiply(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, &multiplyRequest{})
}
