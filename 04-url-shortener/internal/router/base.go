package router

import (
	"net/http"
)

type PageData struct {
	URL string
}

func (r Router) BaseURLGet(w http.ResponseWriter, req *http.Request) {
	r.templator.ExecuteTemplate(w, "base.html", nil)
}

func (r Router) BaseURLPost(w http.ResponseWriter, req *http.Request) {
	url := req.FormValue("url")
	shortURL, _ := r.storage.SaveURL(url)
	r.templator.ExecuteTemplate(w, "shorten.html", PageData{
		URL: shortURL,
	})
}
