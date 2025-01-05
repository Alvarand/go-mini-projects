package handler

import (
	"io"
	"net/http"
)

func Subtract(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "response for subtract\n")
}
