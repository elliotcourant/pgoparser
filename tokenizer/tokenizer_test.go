package tokenizer

import (
	"testing"

	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/stretchr/testify/assert"
)

func testNextToken(t *testing.T, tokenizer *Tokenizer, expected tokens.Token) {
	token, err := tokenizer.nextToken()
	assert.NoError(t, err, "should not have an error while getting next token (%T)", expected)
	assert.Equal(t, expected, token)
}

func TestTokenizer_NextToken(t *testing.T) {
	// Make sure we are consuming whitespace properly.
	t.Run("whitespace", func(t *testing.T) {
		tokenizer := NewTokenizer("\t \n\r\n\r")

		testNextToken(t, tokenizer, common.Tab)
		testNextToken(t, tokenizer, common.Space)
		testNextToken(t, tokenizer, common.Newline)
		testNextToken(t, tokenizer, common.SpecialNewline)
		testNextToken(t, tokenizer, common.Return)

		// EOF
		testNextToken(t, tokenizer, common.EOF)
	})

	t.Run("symbols", func(t *testing.T) {
		tokenizer := NewTokenizer(",;=")

		testNextToken(t, tokenizer, common.Comma)
		testNextToken(t, tokenizer, common.SemiColon)
		testNextToken(t, tokenizer, common.Equals)

		// EOF
		testNextToken(t, tokenizer, common.EOF)
	})
}
