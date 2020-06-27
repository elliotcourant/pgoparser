package symbols

import (
	"github.com/elliotcourant/pgoparser/tokens"
)

var (
	_ Symbol = SymbolType(0)
)

type Symbol interface {
	tokens.Token
	Symbol()
}

//go:generate stringer -type=SymbolType -output=symbols.strings.go
type SymbolType uint8

const (
	_ SymbolType = iota + 100 // try to offset the IDs from the keywords.
	Comma
	Period
	SemiColon
	Equals
	NotEquals
	LessThan
	GreaterThan
	LessThanOrEqualTo
	GreaterThanOrEqualTo
	LeftParentheses
	RightParentheses
	Plus
	Minus
	Division
	Multiply
	Modulo
	Pipe
	StringConcatenation
)

func (_ SymbolType) Token() {}

func (_ SymbolType) Symbol() {}
