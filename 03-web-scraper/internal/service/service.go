package service

import (
	"fmt"
	"net/url"
)

type parser interface {
	Parse()
}

type scraper interface {
	Scrap()
}

type Service interface {
	SetURL(string) error
	Run() error
}

type service struct {
	baseURL string
	parser  parser
	scraper scraper
}

func (s *service) SetURL(URL string) error {
	if _, err := url.ParseRequestURI(URL); err != nil {
		return fmt.Errorf("invalid URL: %s", err)
	}
	s.baseURL = URL
	return nil
}

func (s *service) Run() error {
	return nil

}

func New(parser parser, scraper scraper) Service {
	return &service{
		parser:  parser,
		scraper: scraper,
	}
}
