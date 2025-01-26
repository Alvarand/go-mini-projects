package main

import (
	"errors"
	"fmt"
	"net/http"
	"url-shortener/internal/router"
	"url-shortener/internal/storage/ram"
)

const port = "8080"

func main() {
	storage := ram.New()
	router := router.New(storage)

	mux := http.NewServeMux()

	{
		mux.HandleFunc("POST /", router.BaseURLPost)
		mux.HandleFunc("GET /", router.BaseURLGet)
		mux.HandleFunc("GET /{url}", router.Redirect)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	fmt.Printf("Starting website at localhost:%s\n", port)

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("An error occured:", err)
	}
}
