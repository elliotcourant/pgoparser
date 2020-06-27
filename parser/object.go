package parser

import (
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/words"
)

type ObjectName []Identifier

type Identifier string

func (p *parser) parseObjectName() (ObjectName, error) {
	idents := make(ObjectName, 0)
	for {
		ident, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}

		idents = append(idents, ident)

		if !p.consumeTokenMaybe(tokens.Period{}) {
			break
		}
	}

	return idents, nil
}

func (p *parser) parseIdentifier() (Identifier, error) {
	switch nextToken := p.nextToken().(type) {
	case words.Word:
		return Identifier(nextToken.String()), nil
	default:
		return "", p.expected("identifier", nextToken)
	}
}
