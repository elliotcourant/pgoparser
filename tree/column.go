package tree

import (
	"github.com/elliotcourant/pgoparser/types"
)

type ColumnDefinition struct {
	Name    string
	Type    types.Type
	Options []ColumnOption
}

var (
	_ ColumnOption = PrimaryKey(0)
	_ ColumnOption = Unique(0)
)

type ColumnOption interface {
	ColumnOption()
	String() string
}

type PrimaryKey uint8

func (o PrimaryKey) ColumnOption() {}

func (o PrimaryKey) String() string {
	return "PRIMARY KEY"
}

type Unique uint8

func (o Unique) ColumnOption() {}

func (o Unique) String() string {
	return "UNIQUE"
}

type NotNull uint8

func (o NotNull) ColumnOption() {}

func (o NotNull) String() string {
	return "NOT NULL"
}
