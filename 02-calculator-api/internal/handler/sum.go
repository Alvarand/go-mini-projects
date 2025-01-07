package handler

import (
	"io"
	"net/http"
)

func (h handler) Sum(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "response for sum\n")
}
