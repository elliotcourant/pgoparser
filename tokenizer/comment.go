package tokenizer

import (
	"strings"

	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/pkg/errors"
)

// tokenizeSingleLineComment will consume the current character in the buffer. It will then scan until it reaches the
// end of the buffer, or until a new line is found. tokenizeSingleLineComment should only be called once the second -
// has been peaked and is the current character in the buffer.
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

// tokenizeMultiLineComment will consume the current character in the buffer. It will then scan until it finds */ and
// will return all characters leading up to that. If the EOF is reached before finding the end of the multi line
// comment then an error will be returned.
func (t *Tokenizer) tokenizeMultiLineComment() (tokens.Token, error) {
	var buf strings.Builder
	t.scan() // Consume the * character that should be the current character in the buffer.

	// TODO (elliotcourant) Deal with nested comments. Not sure what these look like, but this is based on the rust
	//  sqlparser.

	maybeClosingComment := false
ScanLoop:
	for {
		character := t.peak()
		switch character {
		case eof:
			// We shouldn't encounter the end of the buffer when we are in a multi-line comment. If we do throw an error.
			return nil, errors.Errorf("unexpected EOF while in multi-line comment")

		default:
			if maybeClosingComment {
				if character == '/' {
					break ScanLoop
				} else {
					// If the character is not / then we need to write *, since we did not write it on the previous loop.
					buf.WriteByte(t.scan()) // Scan the * in.
				}
			}

			// If the current character is a * then we might be closing the comment. Perform the next iteration and if the
			// next character is a / then we can close this comment. If it's not then it will write the * to the buffer
			// there.
			maybeClosingComment = character == '*'

			if !maybeClosingComment {
				buf.WriteByte(t.scan())
			}
		}
	}

	return tokens.MultiLineComment{
		Value: buf.String(),
	}, nil
}
