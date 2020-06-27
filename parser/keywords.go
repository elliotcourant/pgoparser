package parser

import (
	"github.com/elliotcourant/pgoparser/keywords"
)

func (p *parser) parseKeywords(expected ...keywords.Keyword) bool {
	index := p.index
	for _, keyword := range expected {
		if !p.parseKeyword(keyword) {
			p.index = index // If we failed to parse the keyword revert the index back to it's original state.
			return false
		}
	}

	return true
}

func (p *parser) parseKeyword(keyword keywords.Keyword) bool {
	nextToken, index := p.peakTokenIndexed()

	// Type check the keywords just in case.
	if kw, ok := nextToken.(keywords.Keyword); ok && kw == keyword {
		p.index = index // If we were able to parse the keyword then progress the buffer.
		return true
	}

	return false
}
