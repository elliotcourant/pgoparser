package tree

type TypeName struct {
	// Qualified name (list of strings).
	Names []string
	// Is this a set of items.
	SetOf bool
}
