package router

import "net/http"

type PageData struct {
	URL string
}

func (r Router) Redirect(w http.ResponseWriter, req *http.Request) {
	r.templator.ExecuteTemplate(w, "index.html", PageData{
		URL: req.PathValue("url"),
	})
}
