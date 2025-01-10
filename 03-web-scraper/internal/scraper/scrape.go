package scraper

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"time"

	"golang.org/x/net/html"
)

func (s *Scraper) scrape(previousURL string, url string) {
	s.wg.Add(1)
	defer s.wg.Done()

	s.mxParsedURL.Lock()
	if _, isParsed := s.parsedURL[url]; isParsed {
		s.mxParsedURL.Unlock()
		s.mxBadPages.Lock()
		s.mxPagesWithBadLinks.Lock()
		if _, isExists := s.badPages[url]; isExists && !slices.Contains(s.pagesWithBadLinks[previousURL], url) {
			s.pagesWithBadLinks[previousURL] = append(s.pagesWithBadLinks[previousURL], url)
		}
		s.mxBadPages.Unlock()
		s.mxPagesWithBadLinks.Unlock()
		slog.Info(fmt.Sprintf("Page %s is already checked. Skipping", url))
		return
	}

	s.parsedURL[url] = nil
	s.mxParsedURL.Unlock()

	slog.Info(fmt.Sprintf("Checking %s for dead links", url))
	resp, err := http.Get(url)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if resp.StatusCode >= 400 {
		s.mxBadPages.Lock()
		s.badPages[url] = nil
		s.mxBadPages.Unlock()
		s.mxPagesWithBadLinks.Lock()
		if !slices.Contains(s.pagesWithBadLinks[previousURL], url) {
			s.pagesWithBadLinks[previousURL] = append(s.pagesWithBadLinks[previousURL], url)
		}
		s.mxPagesWithBadLinks.Unlock()
		return
	}
	host := resp.Request.URL.Host
	schema := resp.Request.URL.Scheme

	// if host is another -> end
	isSameDomain := host != s.baseHost
	if isSameDomain {
		slog.Warn("Page is not in the same domain. Skipping")
		return
	}

	parsedInfo, err := html.Parse(resp.Body)
	if err != nil {
		slog.Error(err.Error())
	}
	if err := resp.Body.Close(); err != nil {
		slog.Error(err.Error())
	}

	links := getAllLinks([]string{}, parsedInfo)
	for _, link := range links {
		if len(link) == 0 {
			continue
		}
		// link's path starts with /
		if link[0] == '/' {
			link = fmt.Sprintf("%s://%s%s", schema, host, link)
		}
		go s.scrape(url, link)
	}
	time.Sleep(1 * time.Millisecond)
}

func getAllLinks(links []string, node *html.Node) []string {
	if node == nil {
		return links
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = getAllLinks(links, c)
	}

	return links
}
