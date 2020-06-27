package tokens

import (
	"strings"
)

type Identity Token

var (
	_ Token = ObjectName{}
)

type ObjectName []Identity

func NewObjectName(token Token) ObjectName {
	return ObjectName{token}
}

func (o ObjectName) Token() {}

func (o ObjectName) String() string {
	parts := make([]string, len(o))
	for i, part := range o {
		parts[i] = part.String()
	}

	return strings.Join(parts, ".")
}
