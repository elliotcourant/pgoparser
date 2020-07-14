package tree

import (
	"github.com/elliotcourant/pgoparser/types"
)

type ColumnDefinitionOld struct {
	Name     ColumnName
	Type     types.Type
	Nullable bool
	Options  []ColumnOption
}
