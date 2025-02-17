package router

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/exp/slog"
)

var errorGetURL = "failed to get shortURL '%s': %s"

func (r Router) Redirect(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		shortURL := req.PathValue("url")
		url, err := r.storage.GetURL(ctx, shortURL)
		if err != nil {
			slog.Warn(fmt.Sprintf(errorGetURL, shortURL, err))
			w.WriteHeader(http.StatusNotFound)
			r.templator.ExecuteTemplate(w, "not_found.html", nil)
			return
		}
		http.Redirect(w, req, url, http.StatusMovedPermanently)
	}
}
