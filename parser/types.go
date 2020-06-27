package parser

import (
	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/types"
)

func (p *parser) parseDataType() (types.Type, error) {
	switch p.nextToken() {
	case keywords.BOOLEAN:
	case keywords.FLOAT:
	case keywords.REAL:
	case keywords.DOUBLE:
	case keywords.SMALLINT:
		return types.SmallInteger{}, nil
	case keywords.INT, keywords.INTEGER:
	case keywords.BIGINT:
		return types.BigInteger{}, nil
	case keywords.CHAR, keywords.CHARACTER:
	case keywords.UUID:
	case keywords.DATE:
	case keywords.TIMESTAMP:
	case keywords.TIME:
	case keywords.INTERVAL:
	case keywords.REGCLASS:
	case keywords.TEXT:
		return types.Text{}, nil
	case keywords.BYTEA:
	case keywords.NUMERIC, keywords.DECIMAL, keywords.DEC:
	default:
		p.previousToken() // Move the cursor back so we can parse the token as an object name.
		typeName, err := p.parseObjectName()
		if err != nil {
			return nil, err
		}

		return types.Custom{
			Value: typeName.String(),
		}, nil
	}

	return nil, nil
}
