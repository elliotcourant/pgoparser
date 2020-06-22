package quotes

//go:generate stringer --type=Quotes --output=quotes.strings.go
type Quotes uint8

const (
	None Quotes = iota
	Single
	Double
)
