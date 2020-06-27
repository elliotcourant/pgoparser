package tokenizer

import (
	"strings"

	"github.com/elliotcourant/pgoparser/tokens"
)

func (t *Tokenizer) tokenizeSingleLineComment() (tokens.Token, error) {
	var buf strings.Builder
	t.scan() // We want to consume the -. The caller of this method should not have consumed it.

ScanLoop:
	for {
		character := t.peak()
		switch character {
		case '\n', eof:
			break ScanLoop
		default:
			buf.WriteByte(t.scan())
		}
	}

	return tokens.SingleLineComment{
		Value: buf.String(),
	}, nil
}
