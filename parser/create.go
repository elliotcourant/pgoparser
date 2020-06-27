package parser

import (
	"fmt"
	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/tree"
)

func (p *parser) parseCreate() (tree.Statement, error) {
	switch token := p.nextToken().(type) {
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
	ifNotExists := p.parseKeywords(keywords.IF{}, keywords.NOT{}, keywords.EXISTS{})
	fmt.Sprint(ifNotExists)
	return nil, nil
}
