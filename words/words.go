package words

import (
	"github.com/elliotcourant/pgoparser/tokens"
)

type Word interface {
	tokens.Token
	Word()
}
