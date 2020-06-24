package whitespace

import (
	"github.com/elliotcourant/pgoparser/tokens"
)

var (
	_ tokens.Token = Whitespace{}
)

type Whitespace struct {
	Type  Type
	Value string
}

func (w Whitespace) Token() {}

func (w Whitespace) String() string {
	return w.Value
}
