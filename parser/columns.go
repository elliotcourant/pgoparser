package parser

import (
	"fmt"
	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/tree"
	"github.com/elliotcourant/pgoparser/words"
)

func (p *parser) parseColumns() ([]tree.ColumnDefinition, []interface{}, error) {
	// If we cant find the column opening paren then just return?
	if !p.consumeTokenMaybe(tokens.LeftParentheses{}) || p.consumeTokenMaybe(tokens.RightParentheses{}) {
		return nil, nil, nil
	}

	columns := make([]tree.ColumnDefinition, 0)
	constraints := make([]interface{}, 0)

	for {
		// Handle constraints here.
		if constraint, err := p.parseOptionalTableConstraint(); err != nil {
			return nil, nil, err
		} else if constraint != nil {
			constraints = append(constraints, constraint)
		} else if _, ok := p.peakToken().(words.Word); ok { // Try to parse a column name.
			column, err := p.parseColumnDefinition()
			if err != nil {
				return nil, nil, err
			}

			columns = append(columns, column)
		} else {
			return nil, nil, p.expected("column name or constraint definition", p.peakToken())
		}

		comma := p.consumeTokenMaybe(tokens.Comma{})
		if p.consumeTokenMaybe(tokens.RightParentheses{}) {
			break
		} else if !comma {
			return nil, nil, p.expected("',' or ')' after column definition", p.peakToken())
		}
	}

	return columns, constraints, nil
}

func (p *parser) parseOptionalTableConstraint() (_ interface{}, err error) {
	var name string

	// When parsing a constraint, if we grab the constraint keyword first, then the constraint must be named and is
	// specifically declared. If we fail to parse the name or the constraint now then we want to return an error. If not
	// we can still fail gracefully with no error.
	if ok := p.parseKeyword(keywords.CONSTRAINT); ok {
		name, err = p.parseIdentifier()
		if err != nil {
			return nil, err
		}
	}

	token := p.nextToken()
	switch token {
	case keywords.PRIMARY, keywords.UNIQUE:
	case keywords.FOREIGN:
	case keywords.CHECK:
	default:
		if len(name) > 0 {
			return nil, p.expected("PRIMARY, UNIQUE, FOREIGN or CHECK", token)
		} else {
			p.previousToken() // Move the cursor back, we don't need to parse this constraint.
			return nil, nil
		}
	}

	return nil, nil
}

func (p *parser) parseColumnDefinition() (definition tree.ColumnDefinition, err error) {
	definition.Name, err = p.parseIdentifier()
	if err != nil {
		return definition, err
	}

	definition.Type, err = p.parseDataType()
	if err != nil {
		return definition, err
	}

	// If there is a collation defined, parse it.
	if p.parseKeyword(keywords.COLLATE) {
		collation, err := p.parseObjectName()
		if err != nil {
			return definition, err
		}

		// TODO (elliotcourant) Do something with the collation
		fmt.Sprint(collation)
	}

	definition.Options = make([]tree.ColumnOption, 0)
OptionLoop:
	for {
		nextToken := p.peakToken()
		switch nextToken {
		case tokens.EOF{}, tokens.Comma{}, tokens.RightParentheses{}:
			break OptionLoop
		default:
			option, err := p.parseColumnOptionDefinition()
			if err != nil {
				return definition, err
			}

			definition.Options = append(definition.Options, option)
		}
	}

	return definition, nil
}

func (p *parser) parseColumnOptionDefinition() (_ tree.ColumnOption, err error) {
	var name string
	if p.parseKeyword(keywords.CONSTRAINT) {
		name, err = p.parseIdentifier()
		if err != nil {
			return nil, err
		}

		// TODO (elliotcourant) do something with name.
		fmt.Sprint(name)
	}

	nextToken := p.nextToken()
	switch nextToken {
	case keywords.NOT:
		if !p.parseKeyword(keywords.NULL) {
			return nil, p.expected("NULL", p.peakToken())
		}

		return tree.NotNull(0), nil
	case keywords.NULL:
	case keywords.DEFAULT:
	case keywords.PRIMARY:
		if !p.parseKeyword(keywords.KEY) {
			return nil, p.expected("KEY", p.peakToken())
		}

		return tree.PrimaryKey(0), nil
	case keywords.UNIQUE:
		return tree.Unique(0), nil
	case keywords.REFERENCES:
	case keywords.CHECK:

	default:
		return nil, p.expected("column option", nextToken)
	}

	return nil, nil
}
