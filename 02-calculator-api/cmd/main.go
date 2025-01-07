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
		router.HandleFunc("/add", hand.Add)
		router.HandleFunc("/subtract", hand.Subtract)
		router.HandleFunc("/multiply", hand.Multiply)
		router.HandleFunc("/divide", hand.Divide)
		router.HandleFunc("/sum", hand.Sum)
	}

	if err := http.ListenAndServe("localhost:3000", router); err != nil {
		slog.Error(fmt.Sprintf(errorStartServer.Error(), err))
	}
}
