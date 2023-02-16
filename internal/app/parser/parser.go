package parser

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Parser interface {
	Parse() (int, error)
}

type urlParser struct {
	word string
	url  string
}

func New(w string, u string) Parser {
	p := &urlParser{word: w, url: u}
	return p
}

func (p *urlParser) Parse() (int, error) {
	res, err := http.Get(p.url)

	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return 0, err
	}

	count := strings.Count(string(body), p.word)

	return count, nil
}
