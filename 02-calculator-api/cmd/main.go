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

	hand := handler.New()

	{
		router.HandleFunc("POST /add", hand.Add)
		router.HandleFunc("POST /subtract", hand.Subtract)
		router.HandleFunc("POST /multiply", hand.Multiply)
		router.HandleFunc("POST /divide", hand.Divide)
		router.HandleFunc("POST /sum", hand.Sum)
	}

	if err := http.ListenAndServe("localhost:3000", router); err != nil {
		slog.Error(fmt.Sprintf(errorStartServer.Error(), err))
	}
}
