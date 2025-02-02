package router

import (
	"context"
	"html/template"
)

type Storage interface {
	SaveURL(context.Context, string) (string, error)
	GetURL(context.Context, string) (string, error)
}

type Router struct {
	storage   Storage
	templator *template.Template
}

func New(storage Storage) Router {
	return Router{
		storage:   storage,
		templator: template.Must(template.New("").ParseGlob("./templates/*")),
	}
}
