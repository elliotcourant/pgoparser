package tree

type RangeVariable struct {
	// CatalogName is the name of the database, empty if it is not present.
	CatalogName string
	// SchemaName is the name of the schema within the database, empty if it is not present.
	SchemaName string
	// RelationNAme represents the name of the table or the sequence. Should not ever be blank.
	RelationName string

	// Alias is a table alias and eventually optional column aliases.
	Alias *Alias
}
