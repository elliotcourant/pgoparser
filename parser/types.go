package parser

import (
	keywords "github.com/elliotcourant/pgoparser/keywords_v2"
	"github.com/elliotcourant/pgoparser/types"
	"strings"
)

func (p *parser) parseDataType() (types.Type, error) {
	switch p.nextToken() {
	case keywords.BOOLEAN:
	case keywords.FLOAT:
	case keywords.REAL:
	case keywords.DOUBLE:
	case keywords.SMALLINT:
	case keywords.INT, keywords.INTEGER:
	case keywords.BIGINT:
	case keywords.CHAR, keywords.CHARACTER:
	case keywords.UUID:
	case keywords.DATE:
	case keywords.TIMESTAMP:
	case keywords.TIME:
	case keywords.INTERVAL:
	case keywords.REGCLASS:
	case keywords.TEXT:
	case keywords.BYTEA:
	case keywords.NUMERIC, keywords.DECIMAL, keywords.DEC:
	default:
		p.previousToken() // Move the cursor back so we can parse the token as an object name.
		typeName, err := p.parseObjectName()
		if err != nil {
			return nil, err
		}

		return types.Custom{
			Value: strings.Join(typeName, "."),
		}, nil
	}

	return nil, nil
}
