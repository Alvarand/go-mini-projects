package router

import "net/http"

func (r Router) BaseURL(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}
