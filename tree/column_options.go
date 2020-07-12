package tree

import (
	"fmt"
)

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

type Null uint8

func (o Null) ColumnOption() {}

func (o Null) String() string {
	return "NULL"
}

type ForeignKey struct {
	ForeignTable    TableName
	ReferredColumns []ColumnName
	OnDelete        ReferenceAction
	OnUpdate        ReferenceAction
}

func (o ForeignKey) ColumnOption() {}

func (o ForeignKey) String() string {
	// TODO (elliotcourant) Improve stringification of this.
	return fmt.Sprintf("REFERENCES %s (%s)", o.ForeignTable, o.ReferredColumns)
}
