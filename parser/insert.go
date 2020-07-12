package parser

import (
	"fmt"
	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/tree"
)

func (p *parser) parseInsert() (tree.Statement, error) {
	if !p.parseKeyword(keywords.INTO) {
		return nil, p.expected("INTO", p.peakToken())
	}

	targetTable, err := p.parseTableName()
	if err != nil {
		return nil, err
	}

	fmt.Sprint(targetTable)

	cols, err := p.parseParenthesizedColumnList(true)
	if err != nil {
		return nil, err
	}

	fmt.Sprint(cols)

	nextToken := p.nextToken()
	switch nextToken {
	case keywords.VALUES:
		// continue
	case keywords.SELECT:
		// You can insert from a select statement. But I haven't implemented select parsing yet.
		panic("inserting from a select not implemented")
	default:
		return nil, p.expected("VALUES or SELECT", nextToken)
	}

	return nil, nil
}
