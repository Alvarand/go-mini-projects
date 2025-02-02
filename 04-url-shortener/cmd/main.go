package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"url-shortener/internal/client/database/pg"
	"url-shortener/internal/env"
	"url-shortener/internal/router"
	"url-shortener/internal/storage/database"
)

func init() {
	env.Init()
}

func main() {
	ctx := context.Background()

	pgClient, err := pg.New(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create pg client: %s", err))
		return
	}
	defer pgClient.Close(ctx)

	db, err := database.New(ctx, &pgClient)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create database: %s", err))
		return
	}

	router := router.New(db)

	mux := http.NewServeMux()

	{
		mux.HandleFunc("POST /", router.BaseURLPost(ctx))
		mux.HandleFunc("GET /", router.BaseURLGet)
		mux.HandleFunc("GET /{url}", router.Redirect(ctx))
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", env.Get("SERVER_PORT")),
		Handler: mux,
	}

	slog.Info(fmt.Sprintf("Starting website at localhost:%s\n", env.Get("SERVER_PORT")))

	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error(fmt.Sprintf("an error occured: %s", err))
		return
	}
}
