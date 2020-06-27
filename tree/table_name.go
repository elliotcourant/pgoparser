package tree

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type TableName struct {
	CatalogName string
	SchemaName  string
	TableName   string
}

func NewTableName(name []string) (TableName, error) {
	tableName := TableName{
		CatalogName: "",
		SchemaName:  "",
		TableName:   "",
	}

	switch len(name) {
	case 1:
		tableName.TableName = name[0]
	case 2:
		tableName.SchemaName = name[0]
		tableName.TableName = name[1]
	case 3:
		tableName.CatalogName = name[0]
		tableName.SchemaName = name[1]
		tableName.TableName = name[2]
	default:
		return tableName, errors.Errorf("expected 1, 2 or 3 part table name identifier, found %d: ", len(name), name)
	}

	return tableName, nil
}

func (t TableName) String() string {
	var buf strings.Builder
	if len(t.CatalogName) > 0 {
		buf.WriteString(fmt.Sprintf(`"%s".`, t.CatalogName))
	}

	if len(t.SchemaName) > 0 {
		buf.WriteString(fmt.Sprintf(`"%s".`, t.SchemaName))
	}

	buf.WriteString(fmt.Sprintf(`"%s"`, t.TableName))

	return buf.String()
}
