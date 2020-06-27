package parser

import (
	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/tokenizer"
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/tree"
)

type parser struct {
	buffer []tokens.Token
	index  int
}

func Parse(sql string) ([]tree.Statement, error) {
	tokenGenerator := tokenizer.NewTokenizer(sql)
	allTokens, err := tokenGenerator.Tokenize()
	if err != nil {
		return nil, err
	}

	parser := &parser{
		buffer: allTokens,
		index:  0,
	}

	expectingStatementDelimiter := false
	statements := make([]tree.Statement, 0)
	for {
		// Ignore empty statements (between successive statement delimiters)
		for parser.consumeTokenMaybe(tokens.SemiColon{}) {
			expectingStatementDelimiter = false
		}

		nextToken := parser.peakToken()

		// If we really have reached the end of the buffer, then exit gracefully.
		if nextToken == (tokens.EOF{}) {
			break
		}

		if expectingStatementDelimiter {
			return nil, parser.expected("end of statement", nextToken)
		}

		statement, err := parser.parseStatement()
		if err != nil {
			return nil, err
		}

		statements = append(statements, statement)
		expectingStatementDelimiter = true
	}

	return statements, nil
}

func (p *parser) parseStatement() (tree.Statement, error) {
	switch token := p.nextToken().(type) {
	case keywords.SELECT:
		panic("select queries not implemented")
	case keywords.CREATE:
		return p.parseCreate()
	default:
		return nil, p.expected("a sql statement", token)
	}
}
