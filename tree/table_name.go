package tree

import (
	"fmt"
	"strings"
)

type TableName struct {
	CatalogName string
	SchemaName  string
	TableName   string
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
