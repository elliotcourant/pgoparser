package tree

import "github.com/elliotcourant/pgoparser/types"

// ColumnDefinition is a rough translation from PostgreSQL's ColumnDef object. This does have some exclusions from
// PostgreSQL's version.
type ColumnDefinition struct {
	// Name of the column.
	ColumnName string
	// Type of the column.
	Type types.Type
	// Number of times this column is inherited.
	InheritCount int
	// Column has a local (non-inherited) definition.
	IsLocal bool
	// Is the NOT NULL constraint specified. This is just a small helper.
	IsNotNull bool
	// Does this column definition come from a table/composite type.
	IsFromType bool
	// Does this column definition come from a parent partition.
	IsFromParent bool
	// Raw default value from the expression tree. This is not transformed at all.
	RawDefault Expression
	// Default value transformed (should still produce the same value, may be more efficient).
	CookedDefault Expression
	// Column constraints.
	Constraints []Constraint
}
