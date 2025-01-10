package scraper

import (
	"log"
	"sync"
)

type Scraper struct {
	baseURL             string
	baseHost            string
	parsedURL           map[string]interface{}
	badPages            map[string]interface{}
	pagesWithBadLinks   map[string][]string
	mxParsedURL         sync.Mutex
	mxBadPages          sync.Mutex
	mxPagesWithBadLinks sync.Mutex
	wg                  sync.WaitGroup
}

func New(url string) (*Scraper, error) {
	srv := &Scraper{
		parsedURL:         make(map[string]interface{}),
		badPages:          make(map[string]interface{}),
		pagesWithBadLinks: make(map[string][]string),
	}
	err := srv.setURL(url)
	return srv, err
}

func (s *Scraper) Run() error {
	s.scrape("", s.baseURL)

	s.wg.Wait()
	for k, vals := range s.pagesWithBadLinks {
		for _, val := range vals {
			log.Println(k, val)
		}
	}

	return nil

}
