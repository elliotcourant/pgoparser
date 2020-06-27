package tokens

var (
	_ Comment = SingleLineComment{}
	_ Comment = MultiLineComment{}
)

type Comment interface {
	Token
	Comment()
}

type SingleLineComment struct {
	Value string
}

func (s SingleLineComment) Comment() {}

func (s SingleLineComment) Token() {}

func (s SingleLineComment) String() string {
	return s.Value
}

type MultiLineComment struct {
	Value string
}

func (s MultiLineComment) Comment() {}

func (s MultiLineComment) Token() {}

func (s MultiLineComment) String() string {
	return s.Value
}
