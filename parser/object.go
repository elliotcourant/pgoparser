package parser

import (
	"github.com/elliotcourant/pgoparser/symbols"
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

func (p *parser) parseColumnName() (tree.ColumnName, error) {
	objectName, err := p.parseObjectName()
	if err != nil {
		return tree.ColumnName{}, err
	}

	return tree.NewColumnName(objectName)
}

func (p *parser) parseObjectName() (tokens.ObjectName, error) {
	idents := make(tokens.ObjectName, 0)
	for {
		ident, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}

		idents = append(idents, ident)

		if !p.consumeTokenMaybe(symbols.Period) {
			break
		}
	}

	return idents, nil
}

func (p *parser) parseIdentifier() (tokens.Token, error) {
	switch nextToken := p.nextToken().(type) {
	case words.Word:
		return tokens.Identity(nextToken), nil
	default:
		return nil, p.expected("identifier", nextToken)
	}
}
