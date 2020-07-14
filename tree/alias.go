package tree

import "fmt"

var (
	_ Node = Alias{}
)

type Alias struct {
	// Name represents the aliased name for a relation. Never qualified.
	Name string
	// TODO (elliotcourant) Add colum name list.
}

// Node implements the node interface for Alias.
func (a Alias) Node() {}

func (a Alias) String() string {
	// TODO (elliotcourant) Properly escape the name of the alias.
	return fmt.Sprintf("AS %s", a.Name)
}
