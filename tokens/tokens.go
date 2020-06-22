package tokens

type Token interface {
	Token()
	String() string
}

var (
	_ Token = EOF{}
	_ Token = Number{}
	_ Token = Character{}
)

type EOF struct{}

func (E EOF) Token() {}

func (E EOF) String() string {
	return string(byte(0))
}

type Number struct {
	Value string
}

func (n Number) Token() {}

func (n Number) String() string {
	return n.Value
}

type Character struct {
	Value byte
}

func (c Character) Token() {}

func (c Character) String() string {
	return string(c.Value)
}
