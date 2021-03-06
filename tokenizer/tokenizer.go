package tokenizer

import (
	"github.com/elliotcourant/pgoparser/symbols"
	"strings"

	"github.com/elliotcourant/pgoparser/quotes"
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/words"
	"github.com/pkg/errors"
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

func NewTokenizer(str string) *Tokenizer {
	return &Tokenizer{
		input:  str,
		offset: 0,
	}
}

func (t *Tokenizer) Tokenize() ([]tokens.Token, error) {
	// When we create the allTokens array, I guess at the capacity. This will save us a few allocations potentially as
	// I'm basically assuming that every token will be at least 3 bytes worth of data on average.
	allTokens := make([]tokens.Token, 0, len(t.input)/5)

	for {
		token, err := t.nextToken()
		if err != nil {
			return nil, err
		}

		allTokens = append(allTokens, token)

		// If we have reached the end of the buffer then break our loop.
		if _, ok := token.(tokens.EOF); ok {
			break
		}
	}

	return allTokens, nil
}

func (t *Tokenizer) peak() byte {
	if len(t.input) < t.offset+1 {
		return eof
	}

	return t.input[t.offset]
}

func (t *Tokenizer) scan() byte {
	t.offset++
	return t.input[t.offset-1]
}

func (t *Tokenizer) index() int {
	return t.offset
}

// scanAndPeak will consume the current character and will peak the following character.
func (t *Tokenizer) scanAndPeak() byte {
	t.offset++
	return t.peak()
}

func (t *Tokenizer) scanString() string {
	return string(t.scan())
}

func (t *Tokenizer) consumeAndReturn(token tokens.Token) (tokens.Token, error) {
	t.offset++        // Consume the current character.
	return token, nil // And return without an error.
}

func (t *Tokenizer) nextToken() (tokens.Token, error) {
	character := t.peak()
	switch character {
	case eof:
		return common.EOF, nil
	case ' ':
		return t.consumeAndReturn(common.Space)
	case '\t':
		return t.consumeAndReturn(common.Tab)
	case '\r':
		// If the next character is a new line then we want to include that with this character and return it as a single
		// token.
		if nextCharacter := t.scanAndPeak(); nextCharacter == '\n' {
			return t.consumeAndReturn(common.SpecialNewline)
		}

		// If not we can simply return this \r as a token.
		return t.consumeAndReturn(common.Return)
	case '\n':
		return t.consumeAndReturn(common.Newline)
	case '\'':
		// If we encounter an opening single quote, then tokenize a single quoted string.
		return t.tokenizeSingleQuotedString()
	case '"':
		// If we encounter a double quote, then we want to try to tokenize it as a double quoted string.
		return t.tokenizeDoubleQuotedString()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return t.tokenizeNumber()
	case '(':
		return t.consumeAndReturn(symbols.LeftParentheses)
	case ')':
		return t.consumeAndReturn(symbols.RightParentheses)
	case ',':
		return t.consumeAndReturn(symbols.Comma)
	case ';':
		return t.consumeAndReturn(symbols.SemiColon)
	case '=':
		return t.consumeAndReturn(symbols.Equals)
	case '+':
		return t.consumeAndReturn(symbols.Plus)
	case '-':
		// Consume the first - and peak the next character, if the next character is also a - then this is a single line
		// comment.
		if nextCharacter := t.scanAndPeak(); nextCharacter == '-' {
			return t.tokenizeSingleLineComment()
		}

		// We don't want to consume again since we already did a scan above when we peaked.
		return symbols.Minus, nil
	case '/':
		// If the following character is a * then that means this is a multi line comment.
		if nextCharacter := t.scanAndPeak(); nextCharacter == '*' {
			return t.tokenizeMultiLineComment()
		}

		// We don't want to consume again since we already did a scan above when we peaked.
		return symbols.Division, nil
	case '*':
		return t.consumeAndReturn(symbols.Multiply)
	case '%':
		return t.consumeAndReturn(symbols.Modulo)
	case '|':
		if nextCharacter := t.scanAndPeak(); nextCharacter == '|' {
			return t.consumeAndReturn(symbols.StringConcatenation)
		}

		// We don't want to consume again since we already did a scan above when we peaked.
		return symbols.Pipe, nil
	case '.':
		return t.consumeAndReturn(symbols.Period)

	default:
		// If the current character could be the start of an identifier.
		if t.isIdentifierStart(character) {
			// If it is the start, then try to tokenize a word.
			return t.tokenizeWord()
		}
	}

	return nil, nil
}

// tokenizeSingleQuotedString will consume the current starting character (which should be a ') and then read from the
// buffer until it finds another '. It will skip the ' if it is escaped though by a second ' immediately following it.
// It will then return a single quoted word token with a value that is unquoted.
func (t *Tokenizer) tokenizeSingleQuotedString() (tokens.Token, error) {
	var buf strings.Builder
	t.scan() // Consume the first character, this would not have been consumed by the caller but would be '
ScanLoop:
	for {
		character := t.peak()
		switch character {
		case '\'':
			nextCharacter := t.scanAndPeak() // Consume the ' but peak the following character to see if it's escaped.
			// If we find another ' after this one then this one is being escaped and we need to parse it as such.
			if nextCharacter == '\'' {
				// Write a single ' to the string buffer.
				buf.WriteByte(nextCharacter)
			} else {
				// But if the next character is not another ' then that means the ' we saw was the end of the single quoted
				// string.
				break ScanLoop
			}
		case eof:
			// If we reach the end of the file and we have not yet found the closing ' then we should return an error.
			return nil, errors.Errorf("unterminated string literal")
		default:
			// Scan the current byte into the buffer.
			buf.WriteByte(t.scan())
		}
	}

	// Return a single quoted string token.
	return words.NewWord(buf.String(), quotes.Single), nil
}

// tokenizeDoubleQuotedString will consume the current starting character (which should be a ") and then read from the
// buffer until it finds another ". It will not check for any escaped characters as double quoted strings should be
// identifiers in PostgreSQL. If it reaches the end of the file without finding a closing quote then it will return an
// error.
func (t *Tokenizer) tokenizeDoubleQuotedString() (tokens.Token, error) {
	var buf strings.Builder
	t.scan() // Consume the first double quote.

ScanLoop:
	for {
		character := t.peak()
		switch character {
		case '"':
			// We have found another double quote, we can exit our loop.
			break ScanLoop
		case eof:
			// If we reach the end of the file and we have not yet found the closing ' then we should return an error.
			return nil, errors.Errorf("expected close delimiter '\"' before EOF")
		default:
			buf.WriteByte(t.scan())
		}
	}

	// Get the string of the thing we just parsed.
	str := buf.String()

	// Now that we have our string built we need to try to convert it into a word.
	return words.NewWord(str, quotes.Double), nil
}
