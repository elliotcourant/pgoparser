package tokenizer

import (
	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/tokens"
)

// tokenizeWord will consume the first character as the identifier start. It will then consume every valid following
// character until it finds a character that it cannot consume.
func (t *Tokenizer) tokenizeWord() (tokens.Token, error) {
	starting := t.index() // Grab the starting index of the word.

	// Peak the first character, then as long as the current character can be part of an identifier continue to scan and
	// peak the next character while writing to the buffer.
	for character := t.peak(); t.isIdentifierPart(character); character = t.scanAndPeak() {
		// Look forward until we find what would be the end of the word.
	}

	// We now have the entire word. So make a string.
	str := t.input[starting:t.index()]

	// If the word is a valid keyword then use return a new keyword token.
	return keywords.NewKeywordMaybe(str), nil
}
