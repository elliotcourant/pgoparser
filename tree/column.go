package tree

import (
	"fmt"
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/types"
)

type ColumnDefinition struct {
	Name    ColumnName
	Type    types.Type
	Options []ColumnOption
}

type ColumnName struct {
	TableName  TableName
	ColumnName string
}

func NewColumnName(name tokens.ObjectName) (column ColumnName, err error) {
	var tableName TableName
	if len(name) > 1 {
		tableName, err = NewTableName(name[:len(name)-2])
		if err != nil {
			return column, err
		}
	}

	return ColumnName{
		TableName:  tableName,
		ColumnName: name[len(name)-1].String(),
	}, nil
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

type ForeignKey struct {
	ForeignTable    TableName
	ReferredColumns []ColumnName
	OnDelete        interface{}
	OnUpdate        interface{}
}

func (o ForeignKey) ColumnOption() {}

func (o ForeignKey) String() string {
	// TODO (elliotcourant) Improve stringification of this.
	return fmt.Sprintf("REFERENCES %s (%s)", o.ForeignTable, o.ReferredColumns)
}
