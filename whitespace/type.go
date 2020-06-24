package whitespace

//go:generate stringer --type=Type --output=type.strings.go
type Type uint8

const (
	Unknown Type = iota
	Space
	Tab
	Newline
)
