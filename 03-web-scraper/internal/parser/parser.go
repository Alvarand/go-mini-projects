package parser

type Parser interface {
	Parse()
}

type parser struct {
}

func (s parser) Parse() {

}

func New() Parser {
	return parser{}
}
