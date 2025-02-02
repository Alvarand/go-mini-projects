package router

import (
	"context"
	"fmt"
	"net/http"
)

type PageData struct {
	URL string
}

func (r Router) BaseURLGet(w http.ResponseWriter, req *http.Request) {
	r.templator.ExecuteTemplate(w, "base.html", nil)
}

func (r Router) BaseURLPost(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		url := req.FormValue("url")
		shortURL, _ := r.storage.SaveURL(ctx, url)
		r.templator.ExecuteTemplate(w, "shorten.html", PageData{
			URL: fmt.Sprintf("%s/%s", req.Host, shortURL),
		})
	}
}
