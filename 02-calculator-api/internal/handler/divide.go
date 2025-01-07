package handler

import (
	"io"
	"net/http"
)

func (h handler) Divide(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "response for divide\n")
}
