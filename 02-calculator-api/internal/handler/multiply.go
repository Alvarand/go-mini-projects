package handler

import (
	"io"
	"net/http"
)

func Multiply(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "response for multiply\n")
}
