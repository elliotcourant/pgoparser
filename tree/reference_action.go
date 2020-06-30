package tree

//go:generate stringer -type=ReferenceAction -output=reference_action.strings.go
type ReferenceAction uint8

const (
	_ ReferenceAction = iota
	Restrict
	Cascade
	SetNull
	NoAction
	SetDefault
)
