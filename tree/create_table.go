package tree

var (
	_ Statement = CreateTableStatement{}
)

type CreateTableStatement struct {
	// Relation to create.
	Relation    RangeVariable
	IfNotExists bool
	Columns     []ColumnDefinition

	// If the table is a table of a type (like a composite), then that type would be specified here.
	OfType *TypeName
	// TableSpace to use if present, if not present then will be empty.
	TableSpaceName string
}

func (c CreateTableStatement) Statement() {}

func (c CreateTableStatement) Node() {}

func (c CreateTableStatement) String() string {
	return "CREATE TABLE"
}
