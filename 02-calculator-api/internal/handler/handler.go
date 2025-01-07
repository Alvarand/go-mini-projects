package handler

import "net/http"

type handler struct{}

type Handler interface {
	Add(http.ResponseWriter, *http.Request)
	Divide(http.ResponseWriter, *http.Request)
	Multiply(http.ResponseWriter, *http.Request)
	Subtract(http.ResponseWriter, *http.Request)
	Sum(http.ResponseWriter, *http.Request)
}

func New() Handler {
	return handler{}
}
