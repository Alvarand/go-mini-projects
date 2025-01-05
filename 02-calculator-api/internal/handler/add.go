package handler

import (
	"io"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "response for add\n")
}
