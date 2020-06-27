package tokenizer

import (
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/whitespace"
)

type commonTokens struct {
	EOF tokens.EOF

	Space          whitespace.Whitespace
	Tab            whitespace.Whitespace
	Return         whitespace.Whitespace
	SpecialNewline whitespace.Whitespace
	Newline        whitespace.Whitespace

	Comma               tokens.Comma
	Period              tokens.Period
	SemiColon           tokens.SemiColon
	Equals              tokens.Equals
	Plus                tokens.Plus
	Minus               tokens.Minus
	Division            tokens.Division
	Multiply            tokens.Multiply
	LeftParentheses     tokens.LeftParentheses
	RightParentheses    tokens.RightParentheses
	Modulo              tokens.Modulo
	Pipe                tokens.Pipe
	StringConcatenation tokens.StringConcatenation
}

var (
	common = commonTokens{
		EOF: tokens.EOF{},

		Space: whitespace.Whitespace{
			Type:  whitespace.Space,
			Value: " ",
		},
		Tab: whitespace.Whitespace{
			Type:  whitespace.Tab,
			Value: "\t",
		},
		Return: whitespace.Whitespace{
			Type:  whitespace.Newline,
			Value: "\r",
		},
		SpecialNewline: whitespace.Whitespace{
			Type:  whitespace.Newline,
			Value: "\r\n",
		},
		Newline: whitespace.Whitespace{
			Type:  whitespace.Newline,
			Value: "\n",
		},

		Comma:               tokens.Comma{},
		Period:              tokens.Period{},
		SemiColon:           tokens.SemiColon{},
		Equals:              tokens.Equals{},
		Plus:                tokens.Plus{},
		Minus:               tokens.Minus{},
		Division:            tokens.Division{},
		Multiply:            tokens.Multiply{},
		LeftParentheses:     tokens.LeftParentheses{},
		RightParentheses:    tokens.RightParentheses{},
		Modulo:              tokens.Modulo{},
		Pipe:                tokens.Pipe{},
		StringConcatenation: tokens.StringConcatenation{},
	}
)
