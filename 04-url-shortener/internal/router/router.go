package router

import "html/template"

type Storage interface {
	SaveURL(string) (string, error)
	GetURL(string) (string, error)
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
