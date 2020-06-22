package tokens

type Token interface {
	Token()
	String() string
}

var (
	_ Token = EOF{}
	_ Token = Number{}
)

type EOF struct{}

func (E EOF) Token() {
	panic("implement me")
}

func (E EOF) String() string {
	return string(byte(0))
}

type Number struct {
	Value string
}

func (n Number) Token() {
	panic("implement me")
}

func (n Number) String() string {
	return n.Value
}
