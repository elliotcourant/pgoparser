package tokenizer

import (
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/whitespace"
)

var (
	eof = byte(0)
)

type (
	Tokenizer struct {
		input  string
		offset int
	}
)

func (t *Tokenizer) peak() byte {
	if len(t.input) < t.offset+1 {
		return eof
	}

	return t.input[t.offset]
}

func (t *Tokenizer) scan() byte {
	if len(t.input) < t.offset+1 {
		return eof
	}

	t.offset++

	return t.input[t.offset-1]
}

func (t *Tokenizer) scanString() string {
	return string(t.scan())
}

func (t *Tokenizer) nextToken() tokens.Token {
	character := t.peak()
	switch character {
	case ' ':
		return whitespace.Whitespace{
			Type:  whitespace.Space,
			Value: t.scanString(),
		}
	case '\t':
		return whitespace.Whitespace{
			Type:  whitespace.Tab,
			Value: t.scanString(),
		}
	case '\n', '\r':
		// TODO (elliotcourant) Add handling of the return.
		return whitespace.Whitespace{
			Type:  whitespace.Newline,
			Value: t.scanString(),
		}
	case '\'':
		panic("single quoted string not implemented")
	case '"':
		panic("double quoted string not implemented")
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		panic("numbers not implemented")
	case '(':

	case ',':
		return tokens.Comma{}
	case ';':
		return tokens.SemiColon{}
	}

	return nil
}
