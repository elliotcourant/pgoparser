package words

import (
	"github.com/elliotcourant/pgoparser/quotes"
	"github.com/elliotcourant/pgoparser/tokens"
)

type Word interface {
	tokens.Token
	Word()
	Quotes() quotes.Quotes
}

type base struct {
	Value     string
	QuoteType quotes.Quotes
}

func (b base) Quotes() quotes.Quotes {
	return b.QuoteType
}

func (b base) Token() {}

func (b base) String() string {
	return b.Value
}

func (b base) Word() {}

func NewWord(value string, quoteType quotes.Quotes) Word {
	return &base{
		Value:     value,
		QuoteType: quoteType,
	}
}
