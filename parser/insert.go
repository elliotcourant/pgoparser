package parser

import (
	"fmt"
	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/symbols"
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/tree"
	"github.com/elliotcourant/pgoparser/words"
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
		rows := make([][]tokens.Token, 0)
		_, err := p.parseCommaSeparated(func() (tokens.Token, error) {
			if !p.consumeTokenMaybe(symbols.LeftParentheses) {
				return nil, p.expected("(", p.peakToken())
			}
			values, err := p.parseCommaSeparated(func() (tokens.Token, error) {
				switch nextToken := p.nextToken().(type) {
				case words.Word, tokens.Number:
					return nextToken, nil
				default:
					return nil, p.expected("string or integer", nextToken)
				}
			})
			if err != nil {
				return nil, err
			}
			if !p.consumeTokenMaybe(symbols.RightParentheses) {
				return nil, p.expected(")", p.peakToken())
			}

			rows = append(rows, values)
			return nil, nil
		})

		if err != nil {
			return nil, err
		}

		fmt.Sprint(rows)
		// continue
	case keywords.SELECT:
		// You can insert from a select statement. But I haven't implemented select parsing yet.
		panic("inserting from a select not implemented")
	default:
		return nil, p.expected("VALUES or SELECT", nextToken)
	}

	return nil, nil
}
