package tokens

type SingleLineComment struct {
	Value string
}

func (s SingleLineComment) Token() {}

func (s SingleLineComment) String() string {
	return s.Value
}

type MultiLineComment struct {
	Value string
}

func (s MultiLineComment) Token() {}

func (s MultiLineComment) String() string {
	return s.Value
}
