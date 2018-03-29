package engine

import (
	"log"

	"github.com/bean-du/crawler/fatcher"
)

func worker(r Request) (ParseResult, error) {
	log.Printf("fatching Url: %s", r.Url)
	body, err := fatcher.Fatch(r.Url)
	if err != nil {
		log.Printf("Fatcher: error fatching url %s : %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
