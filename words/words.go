package words

import (
	"github.com/elliotcourant/pgoparser/quotes"
	"github.com/elliotcourant/pgoparser/tokens"
)

var (
	_ Word = String{}
	_ Word = SingleQuotedString{}
	_ Word = DoubleQuotedString{}
)

type Word interface {
	tokens.Token
	Word()
	Quotes() quotes.Quotes
}

func NewWord(value string, quoteType quotes.Quotes) Word {
	switch quoteType {
	case quotes.None:
		return String{
			Value: value,
		}
	case quotes.Single:
		return SingleQuotedString{
			Value: value,
		}
	case quotes.Double:
		return DoubleQuotedString{
			Value: value,
		}
	default:
		panic("invalid quote type")
	}
}

type String struct {
	Value string
}

func (b String) Quotes() quotes.Quotes {
	return quotes.None
}

func (b String) Token() {}

func (b String) String() string {
	return b.Value
}

func (b String) Word() {}

type SingleQuotedString struct {
	Value string
}

func (s SingleQuotedString) Quotes() quotes.Quotes {
	return quotes.Single
}

func (s SingleQuotedString) Token() {}

func (s SingleQuotedString) String() string {
	return s.Value
}

func (s SingleQuotedString) Word() {}

type DoubleQuotedString struct {
	Value string
}

func (s DoubleQuotedString) Quotes() quotes.Quotes {
	return quotes.Single
}

func (s DoubleQuotedString) Token() {}

func (s DoubleQuotedString) String() string {
	return s.Value
}

func (s DoubleQuotedString) Word() {}
