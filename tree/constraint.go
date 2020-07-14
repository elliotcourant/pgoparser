package tree

type Constraint struct {
	Type ConstraintType

	/*
		Fields used for most or all constraint types.
	*/

	// Name is the constraint name if it's present. If it is not then this will be nil.
	Name *string
	// Is the constraint deferrable or not? Defaults to false.
	Deferrable bool
	// Is the constraint initally deferred.
	InitiallyDeferred bool

	/*
		Fields used for constraints with expressions; CHECK and DEFAULT constraints.
	*/

	// Is the constraint non-inheritable.
	IsNoInherit bool
	// Expression as an untransformed tree.
	RawExpression Expression
	// Expression as a string representation, is nil if no expression is present.
	CookedExpression *string
}
