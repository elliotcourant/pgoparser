package parser

import (
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/tree"
	"github.com/elliotcourant/pgoparser/words"
)

func (p *parser) parseTableName() (tree.TableName, error) {
	objectName, err := p.parseObjectName()
	if err != nil {
		return tree.TableName{}, err
	}

	return tree.NewTableName(objectName)
}

func (p *parser) parseObjectName() ([]string, error) {
	idents := make([]string, 0)
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

func (p *parser) parseIdentifier() (string, error) {
	switch nextToken := p.nextToken().(type) {
	case words.Word:
		return nextToken.String(), nil
	default:
		return "", p.expected("identifier", nextToken)
	}
}
