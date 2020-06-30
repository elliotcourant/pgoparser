package tree

import (
	"github.com/elliotcourant/pgoparser/types"
)

type ColumnDefinition struct {
	Name    ColumnName
	Type    types.Type
	Options []ColumnOption
}
