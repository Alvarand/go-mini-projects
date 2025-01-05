package main

import (
	"calculator-api/internal/handler"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

var errorStartServer = errors.New("failed to start server: %s")

func main() {

	router := http.NewServeMux()

	{
		router.HandleFunc("/add", handler.Add)
		router.HandleFunc("/subtract", handler.Subtract)
		router.HandleFunc("/multiply", handler.Multiply)
		router.HandleFunc("/divide", handler.Divide)
		router.HandleFunc("/sum", handler.Sum)
	}

	if err := http.ListenAndServe("localhost:3000", router); err != nil {
		slog.Error(fmt.Sprintf(errorStartServer.Error(), err))
	}
}
