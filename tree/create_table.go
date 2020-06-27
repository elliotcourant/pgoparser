package tree

var (
	_ Statement = CreateTableStatement{}
)

type CreateTableStatement struct {
	IfNotExists bool
	TableName   TableName
	Columns     []ColumnDefinition
}

func (c CreateTableStatement) Statement() {}

func (c CreateTableStatement) String() string {
	return "CREATE TABLE"
}
