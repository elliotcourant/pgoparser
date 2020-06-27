package tokenizer

import (
	"strings"

	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/quotes"
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/words"
)

// tokenizeWord will consume the first character as the identifer start. It will then consume every valid following
// character until it finds a character that it cannot consume.
func (t *Tokenizer) tokenizeWord() (tokens.Token, error) {
	var buf strings.Builder
	buf.WriteByte(t.scan()) // Consume the first byte, since we know it will be an identifier start.

	// Peak the first character, then as long as the current character can be part of an identifier continue to scan and
	// peak the next character while writing to the buffer.
	for character := t.peak(); t.isIdentifierPart(character); character = t.scanAndPeak() {
		buf.WriteByte(character)
	}

	// We now have the entire word. So make a string.
	str := buf.String()

	// If the word is a valid keyword then use return a new keyword token.
	if keywords.IsValidKeyword(str) {
		return keywords.NewKeyword(str), nil
	}

	// Otherwise just return a generic word token.
	return words.NewWord(str, quotes.None), nil
}
