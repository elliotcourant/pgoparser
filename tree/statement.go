package tree

type (
	// Node represents any part of the syntax tree.
	Node interface {
		Node()
		String() string
	}

	// Statement represents nodes that are top level objects within the syntax tree. These are wrappers around the entire
	// user's input.
	Statement interface {
		Statement()
		Node
	}

	// Expressions are nodes that represent one small portion of a statement. This includes comparisons or operations on
	// references or objects within a statement.
	Expression interface {
		Expression()
		Node
	}
)
