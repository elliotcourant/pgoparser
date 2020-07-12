package parser

import (
	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/symbols"
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
		for parser.consumeTokenMaybe(symbols.SemiColon) {
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
	token := p.nextToken()
	switch token {
	case keywords.SELECT:
		panic("select queries not implemented")
	case keywords.CREATE:
		return p.parseCreate()
	case keywords.INSERT:
		return p.parseInsert()
	default:
		return nil, p.expected("a sql statement", token)
	}
}

func (p *parser) expectToken(token tokens.Token) error {
	if p.consumeTokenMaybe(token) {
		return nil
	} else {
		return p.expected(token, p.peakToken())
	}
}

func (p *parser) parseCommaSeparated(do func() (tokens.Token, error)) ([]tokens.Token, error) {
	values := make([]tokens.Token, 0)
	for {
		token, err := do()
		if err != nil {
			return nil, err
		}

		values = append(values, token)

		if !p.consumeTokenMaybe(symbols.Comma) {
			break // If the next token is not a comma then break the loop.
		}
	}

	return values, nil
}
