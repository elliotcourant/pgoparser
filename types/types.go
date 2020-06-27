package types

type Type interface {
	Type()
	String() string
}

type Custom struct {
	Value string
}

func (c Custom) Type() {}

func (c Custom) String() string {
	return c.Value
}

type BigInt struct{}

func (b BigInt) Type() {}

func (b BigInt) String() string {
	return "BIGINT"
}
