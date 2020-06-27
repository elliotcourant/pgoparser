package parser

import (
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/tree"
	"github.com/elliotcourant/pgoparser/words"
)

func (p *parser) parseColumns() ([]tree.ColumnDefinition, []interface{}, error) {
	columns := make([]tree.ColumnDefinition, 0)
	constraints := make([]interface{}, 0)

	// If we cant find the column opening paren then just return?
	if !p.consumeTokenMaybe(tokens.LeftParentheses{}) || p.consumeTokenMaybe(tokens.RightParentheses{}) {
		return columns, constraints, nil
	}

	for {
		// Handle constraints here.

		// If the next token is a word, then we are parsing a column definition.
		if _, ok := p.peakToken().(words.Word); ok {

		}
	}
}

func (p *parser) parseColumnDefinition() (definition tree.ColumnDefinition, _ error) {
	name, err := p.parseIdentifier()
	if err != nil {
		return definition, err
	}

}
