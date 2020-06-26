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

// scanAndPeak will consume the current character and will peak the following character.
func (t *Tokenizer) scanAndPeak() byte {
	t.scan()
	return t.peak()
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
	case '\r':
		// If the next character is a new line then we want to include that with this character and return it as a single
		// token.
		if nextCharacter := t.scanAndPeak(); nextCharacter == '\n' {
			return whitespace.Whitespace{
				Type:  whitespace.Newline,
				Value: "\r\n",
			}
		}

		// If not we can simply return this \r as a token.
		return whitespace.Whitespace{
			Type:  whitespace.Newline,
			Value: "\r",
		}
	case '\n':
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
	case '=':
		return tokens.Equals{}
	}

	return nil
}
