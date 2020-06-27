package tokenizer

import (
	"github.com/elliotcourant/pgoparser/symbols"
	"testing"

	"github.com/elliotcourant/pgoparser/keywords"
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/words"
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

	t.Run("single quoted string", func(t *testing.T) {
		tokenizer := NewTokenizer(`'this is a single quoted string'`)

		token, err := tokenizer.nextToken()
		assert.NoError(t, err, "should not have an error single quoted string")
		assert.Implements(t, new(words.Word), token)
		assert.Equal(t, "this is a single quoted string", token.String())
	})

	t.Run("double quoted string", func(t *testing.T) {
		tokenizer := NewTokenizer(`"this is a double quoted string"`)

		token, err := tokenizer.nextToken()
		assert.NoError(t, err, "should not have an error double quoted string")
		assert.Implements(t, new(words.Word), token)
		assert.Equal(t, "this is a double quoted string", token.String())
	})

	t.Run("symbols", func(t *testing.T) {
		// TODO (elliotcourant) Add tests for more symbols.
		tokenizer := NewTokenizer(",;=()+-")

		testNextToken(t, tokenizer, symbols.Comma)
		testNextToken(t, tokenizer, symbols.SemiColon)
		testNextToken(t, tokenizer, symbols.Equals)
		testNextToken(t, tokenizer, symbols.LeftParentheses)
		testNextToken(t, tokenizer, symbols.RightParentheses)
		testNextToken(t, tokenizer, symbols.Plus)
		testNextToken(t, tokenizer, symbols.Minus)

		// EOF
		testNextToken(t, tokenizer, common.EOF)
	})

	t.Run("numbers", func(t *testing.T) {
		str := "1234.56"
		tokenizer := NewTokenizer(str)

		token, err := tokenizer.nextToken()
		assert.NoError(t, err, "should not have an error parsing number")
		assert.IsType(t, tokens.Number{}, token)
		assert.Equal(t, str, token.String())
	})

	t.Run("single line comment", func(t *testing.T) {
		text := `--this is a test`

		tokenizer := NewTokenizer(text)

		token, err := tokenizer.nextToken()
		assert.NoError(t, err, "should not have an error parsing single line comment")
		assert.IsType(t, tokens.SingleLineComment{}, token)
		assert.Equal(t, "this is a test", token.String())
	})

	t.Run("multi line comment", func(t *testing.T) {
		text := `/* i am
		a multi line
		commen */`

		tokenizer := NewTokenizer(text)

		token, err := tokenizer.nextToken()
		assert.NoError(t, err, "should not have an error parsing multi line comment")
		assert.IsType(t, tokens.MultiLineComment{}, token)
	})

	t.Run("keyword", func(t *testing.T) {
		tokenizer := NewTokenizer("SELECT things FROM stuff")

		token, err := tokenizer.nextToken()
		assert.NoError(t, err)
		assert.IsType(t, keywords.SELECT, token)

		testNextToken(t, tokenizer, common.Space)

		token, err = tokenizer.nextToken()
		assert.NoError(t, err)
		assert.IsType(t, words.String{}, token)
		assert.Equal(t, "things", token.String())

		testNextToken(t, tokenizer, common.Space)

		token, err = tokenizer.nextToken()
		assert.NoError(t, err)
		assert.IsType(t, keywords.FROM, token)

		testNextToken(t, tokenizer, common.Space)

		token, err = tokenizer.nextToken()
		assert.NoError(t, err)
		assert.IsType(t, words.String{}, token)
		assert.Equal(t, "stuff", token.String())
	})
}
