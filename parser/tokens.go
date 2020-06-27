package parser

import (
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/whitespace"
)

// consumeTokenMaybe will return true if the next non-whitespace token matches the provided token. It will then move the
// current index to that position. If the next non-whitespace token does not match it will return false and do nothing.
func (p *parser) consumeTokenMaybe(token tokens.Token) bool {
	if nextToken, index := p.peakTokenIndexed(); nextToken == token {
		p.index = index // Move the index forward
		return true
	}

	return false
}

// nextToken peaks the next non-whitespace token and then updates the current index to the index of that token and
// returns its value.
func (p *parser) nextToken() (token tokens.Token) {
	token, p.index = p.peakNthToken(0)
	return
}

// previousToken will move the cursor back to the last non-whitespace/non-comment token.
func (p *parser) previousToken() tokens.Token {
	index := p.index
	for {
		index -= 1

		// Prevent an out of range error.
		if index > len(p.buffer)-1 {
			return tokens.EOF{}
		}

		token := p.buffer[index-1]

		switch token.(type) {
		case whitespace.Whitespace, tokens.Comment:
			// If the token is whitespace or a comment skip it.
			continue
		default:
			p.index = index
			return token
		}
	}
}

// peakToken will return the first non-whitespace token that has not yet been processed. Or it will return the EOF.
func (p *parser) peakToken() tokens.Token {
	token, _ := p.peakTokenIndexed()
	return token
}

// peakTokenIndexed will return the first non-whitespace token that has not yet been processed. Or it will return the
// EOF. It will also return the new index if you decide to move forward to the specified token.
func (p *parser) peakTokenIndexed() (tokens.Token, int) {
	return p.peakNthToken(0)
}

// return the nth non-whitespace/comment token that has not yet been processed.
func (p *parser) peakNthToken(n int) (token tokens.Token, index int) {
	x := 0
	index = p.index
	for {
		index += 1

		// Prevent an out of range error.
		if index > len(p.buffer)-1 {
			return tokens.EOF{}, index
		}

		token = p.buffer[index-1]

		switch token.(type) {
		case whitespace.Whitespace, tokens.Comment:
			// If the token is whitespace or a comment skip it.
			continue
		default:
			if x == n {
				return token, index
			}

			// Count up until we get the index we want.
			x++
		}
	}
}
