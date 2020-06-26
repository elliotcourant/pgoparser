package tokenizer

import (
	"strings"

	"github.com/elliotcourant/pgoparser/tokens"
)

// tokenizeNumber will consume the current character (which should be a number) and will consume until a non-number and
// non-period character is encountered.
func (t *Tokenizer) tokenizeNumber() (tokens.Token, error) {
	var buf strings.Builder
	buf.WriteByte(t.scan()) // Consume the current character, which should be a number.

ScanLoop:
	for {
		character := t.peak()
		switch character {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			buf.WriteByte(t.scan())
		default:
			break ScanLoop
		}
	}

	return tokens.Number{
		Value: buf.String(),
	}, nil
}
