package parser

import (
	"github.com/elliotcourant/pgoparser/symbols"
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/tree"
	"github.com/elliotcourant/pgoparser/words"
	"github.com/pkg/errors"
)

// parseTableName will parse the words that are period delimited until the end is reached. It will return
// a range variable with the table reference. It does not currently parse aliases.
func (p *parser) parseTableName() (tree.RangeVariable, error) {
	objectName, err := p.parseObjectName()
	if err != nil {
		return tree.RangeVariable{}, err
	}

	var rangeVar tree.RangeVariable
	switch len(objectName) {
	case 1:
		rangeVar.RelationName = objectName[0].String()
	case 2:
		rangeVar.SchemaName = objectName[0].String()
		rangeVar.RelationName = objectName[1].String()
	case 3:
		rangeVar.CatalogName = objectName[0].String()
		rangeVar.SchemaName = objectName[1].String()
		rangeVar.RelationName = objectName[2].String()
	default:
		return tree.RangeVariable{}, errors.Errorf("unexpect object name length: %d", len(objectName))
	}

	// TODO (elliotcourant) Add support for aliases maybe?

	return rangeVar, nil
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
