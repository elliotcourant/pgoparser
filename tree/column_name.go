package tree

import (
	"github.com/elliotcourant/pgoparser/tokens"
)

type ColumnName struct {
	TableName  string
	ColumnName string
}

func NewColumnName(name tokens.ObjectName) (column ColumnName, err error) {
	var tableName string
	if len(name) > 1 {
		tableName = name[0].String()
	}

	return ColumnName{
		TableName:  tableName,
		ColumnName: name[len(name)-1].String(),
	}, nil
}
