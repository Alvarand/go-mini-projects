package router

import (
	"fmt"
	"net/http"

	"golang.org/x/exp/slog"
)

func (r Router) Redirect(w http.ResponseWriter, req *http.Request) {
	shortURL := req.PathValue("url")
	url, err := r.storage.GetURL(shortURL)
	if err != nil {
		slog.Warn(fmt.Sprintf("failed to get shortURL '%s': %s", shortURL, err))
		w.WriteHeader(http.StatusNotFound)
		r.templator.ExecuteTemplate(w, "not_found.html", nil)
		return
	}
	http.Redirect(w, req, url, http.StatusMovedPermanently)
}
