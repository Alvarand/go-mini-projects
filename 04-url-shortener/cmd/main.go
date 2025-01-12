package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	slog.Info("Starting website at localhost:8080")

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error(fmt.Sprint("An error occured:", err))
	}
}
