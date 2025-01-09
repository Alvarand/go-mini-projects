package main

import (
	"flag"
	"fmt"
	"log/slog"
	"web-scraper/internal/scraper"
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

	service, err := scraper.New(baseURL)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create service: %s", err))
		return
	}

	if err := service.Run(); err != nil {
		slog.Error(fmt.Sprintf("failed to run service: %s", err))
		return
	}
}
