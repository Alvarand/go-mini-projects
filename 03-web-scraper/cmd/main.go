package main

import (
	"flag"
	"fmt"
	"log/slog"
	"web-scraper/internal/parser"
	"web-scraper/internal/scraper"
	"web-scraper/internal/service"
)

var baseURL = ""
var errorEmptyURL = "url is empty"

func init() {
	flag.StringVar(&baseURL, "url", "", "input URL to parse and scrap")
	flag.Parse()
}

func main() {
	if baseURL == "" {
		slog.Error(errorEmptyURL)
		return
	}

	parser := parser.New()
	scraper := scraper.New()

	service := service.New(parser, scraper)
	if err := service.SetURL(baseURL); err != nil {
		slog.Error(fmt.Sprintf("failed to set URL: %s", err))
		return
	}

	if err := service.Run(); err != nil {
		slog.Error(fmt.Sprintf("failed to run service: %s", err))
		return
	}
}
