package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"url-shortener/internal/client/database/pg"
	"url-shortener/internal/env"
	"url-shortener/internal/router"
	"url-shortener/internal/storage/ram"
)

func init() {
	env.Init()
}

func main() {
	ctx := context.Background()

	pgClient, err := pg.New(ctx)
	if err != nil {
		log.Fatalf("failed to create pg client: %s", err)
	}
	defer pgClient.Close(ctx)

	storage := ram.New()
	router := router.New(storage)

	mux := http.NewServeMux()

	{
		mux.HandleFunc("POST /", router.BaseURLPost)
		mux.HandleFunc("GET /", router.BaseURLGet)
		mux.HandleFunc("GET /{url}", router.Redirect)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", env.Get("SERVER_PORT")),
		Handler: mux,
	}

	fmt.Printf("Starting website at localhost:%s\n", env.Get("SERVER_PORT"))

	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("An error occured:", err)
	}
}
