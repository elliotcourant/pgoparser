package parser

import (
	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/tree"
)

func (p *parser) parseCreate() (tree.Statement, error) {
	token := p.nextToken()
	switch token {
	case keywords.TABLE:
		return p.parseCreateTable()
	case keywords.VIEW:
		panic("not implemented")
	case keywords.INDEX:
		panic("not implemented")
	case keywords.SCHEMA:
		panic("not implemented")
	default:
		return nil, p.expected("TABLE, VIEW, INDEX or SCHEMA after CREATE", token)
	}
}

func (p *parser) parseCreateTable() (tree.Statement, error) {
	ifNotExists := p.parseKeywords(keywords.IF, keywords.NOT, keywords.EXISTS)
	tableName, err := p.parseTableName()
	if err != nil {
		return nil, err
	}

	columns, _, err := p.parseColumns()
	if err != nil {
		return nil, err
	}

	return tree.CreateTableStatement{
		IfNotExists: ifNotExists,
		TableName:   tableName,
		Columns:     columns,
	}, nil
}
