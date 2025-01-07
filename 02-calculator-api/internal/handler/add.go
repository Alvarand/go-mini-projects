package handler

import (
	"io"
	"net/http"
)

func (h handler) Add(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "response for add\n")
}
