package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
)

type parser struct {
	dom *goquery.Document
}

func InitDom(html string) (*parser, error) {

	p := new(parser)

	htmlBuf := bytes.NewBufferString(html)
	dom, err := goquery.NewDocumentFromReader(htmlBuf)
	if err != nil {

		return p, err
	}

	p.dom = dom
	return p, nil
}

func (query *parser) Find(selector string) *goquery.Selection {
	return query.dom.Find(selector)
}

func (query *parser) FindAll(selector string) []*goquery.Selection {
	var doms []*goquery.Selection
	query.dom.Find(selector).Each(func(i int, s *goquery.Selection) {
		doms = append(doms, s)
	})
	return doms
}
