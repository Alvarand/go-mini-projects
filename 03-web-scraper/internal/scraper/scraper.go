package scraper

type Scraper interface {
	Scrap()
}

type scraper struct {
}

func (s scraper) Scrap() {

}

func New() Scraper {
	return scraper{}
}
