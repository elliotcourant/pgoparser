package tree

type ConstraintType uint8

const (
	// Not standard SQL but a lot of people expect it.
	ConstraintTypeNull ConstraintType = iota
	ConstraintTypeNotNull
	ConstraintTypeDefault
	ConstraintTypeIdentity
	ConstraintTypeCheck
	ConstraintTypePrimary
	ConstraintTypeUnique
	ConstraintTypeExclusion
	ConstraintTypeForeign

	// Attributes for the previous constraint node.
	ConstraintTypeAttributeDeferrable
	ConstraintTypeAttributeNotDeferrable
	ConstraintTypeAttributeDeferred
	ConstraintTypeAttributeImmediate
)
