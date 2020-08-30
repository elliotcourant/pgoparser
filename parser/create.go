package parser

import (
	. "github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/symbols"
	"github.com/elliotcourant/pgoparser/tree"
)

func (p *parser) parseCreateOrReplace() (tree.Statement, error) {
	token := p.nextToken()
	switch token {
	case AGGREGATE:
		panic("aggregates not implemented")
	default:
		return nil, p.expected("AGGREGATE or VIEW after CREATE OR REPLACE", token)
	}
}

func (p *parser) parseCreate() (tree.Statement, error) {
	// TODO (elliotcourant) Add support for global|local.

	// token := p.nextToken()
	// switch token {
	// case ACCESS:
	// 	if !p.parseKeyword(METHOD) {
	// 		return nil, p.expected("METHOD after ACCESS", token)
	// 	}
	// 	panic("access method not implemented")
	// case CAST:
	// 	panic("cast not implemented")
	// case COLLATION:
	// 	panic("collation not implemented")
	// case DEFAULT:
	// 	// Currently i think this is only for conversions.
	// 	token = p.nextToken()
	// 	switch token {
	// 	case CONVERSION:
	// 		panic("conversion not implemented")
	// 	default:
	// 		return nil, p.expected("CONVERSION after CREATE DEFAULT", token)
	// 	}
	// case CONVERSION:
	// 	panic("conversion not implemented")
	// case DATABASE:
	// 	panic("database not implemented")
	// case DOMAIN:
	// 	panic("domain not implemented")
	// case EVENT:
	// 	if !p.parseKeyword(TRIGGER) {
	// 		return nil, p.expected("TRIGGER after CREATE EVENT", token)
	// 	}
	// 	panic("event trigger not implemented")
	// case EXTENSION:
	// 	panic("extension not implemented")
	// case FOREIGN:
	// 	token = p.nextToken()
	// 	switch token {
	// 	case DATA:
	// 		if !p.parseKeyword(WRAPPER) {
	// 			return nil, p.expected("WRAPPER after CREATE FOREIGN DATA", token)
	// 		}
	// 		panic("foreign data wrapper not implemented")
	// 	case TABLE:
	// 		panic("foreign table not implemented")
	// 	default:
	// 		return nil, p.expected("DATA WRAPPER or TABLE after CREATE FOREIGN", token)
	//
	// 	}
	// case FUNCTION:
	// 	panic("function not implemented")
	// }

	var temporary, unlogged bool

	// If the next keyword is temp/temporary then this is a temp table.
	temporary = p.parseAnyKeyword(TEMP, TEMPORARY)
	unlogged = p.parseKeyword(UNLOGGED)

	token := p.nextToken()
	switch token {
	case AGGREGATE:
		if temporary {
			return nil, p.expected("VIEW or TABLE", token)
		} else if unlogged {
			return nil, p.expected("TABLE", token)
		}
		panic("aggregate not implemented")
	case TABLE:
		return p.parseCreateTable(temporary, unlogged)
	case RECURSIVE:
		if !p.parseKeyword(VIEW) {
			return nil, p.expected("VIEW", token)
		}

		// If we get a recursive keyword, then fallthrough to a view.
		fallthrough
	case VIEW:
		panic("not implemented")
	case INDEX:
		panic("not implemented")
	case SCHEMA:
		panic("not implemented")
	default:
		return nil, p.expected("TABLE, VIEW, INDEX or SCHEMA after CREATE", token)
	}
}

func (p *parser) parseCreateTable(temporary, unlogged bool) (tree.Statement, error) {
	ifNotExists := p.parseKeywords(IF, NOT, EXISTS)

	tableName, err := p.parseTableName()
	if err != nil {
		return nil, err
	}

	columns, _, err := p.parseColumns()
	if err != nil {
		return nil, err
	}

	inherits := make([]tree.RangeVariable, 0)
	if p.parseKeyword(INHERITS) {
		for {
			table, err := p.parseTableName()
			if err != nil {
				return nil, err
			}

			inherits = append(inherits, table)

			if !p.consumeTokenMaybe(symbols.Comma) {
				break
			}
		}
	}

	return tree.CreateTableStatement{
		IfNotExists: ifNotExists,
		Temporary:   temporary,
		Unlogged:    unlogged,
		Relation:    tableName,
		Columns:     columns,
		Inherits:    inherits,
	}, nil
}
