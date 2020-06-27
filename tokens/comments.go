package tokens

type SingleLineComment struct {
	Value string
}

func (s SingleLineComment) Token() {}

func (s SingleLineComment) String() string {
	return s.Value
}
