package scraper

import (
	"fmt"
	"net/url"
)

func (s *Scraper) setURL(URL string) error {
	parsedURL, err := url.ParseRequestURI(URL)
	if err != nil {
		return fmt.Errorf("invalid URL: %s", err)
	}

	s.baseURL = URL
	s.baseHost = parsedURL.Host
	return nil
}
