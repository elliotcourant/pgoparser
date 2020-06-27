package types

type Type interface {
	Type()
	String() string
}

type Custom struct {
	Value string
}

func (t Custom) Type() {}

func (t Custom) String() string {
	return t.Value
}

type SmallInteger struct{}

func (t SmallInteger) Type() {}

func (t SmallInteger) String() string {
	return "SMALLINT"
}

type BigInteger struct{}

func (t BigInteger) Type() {}

func (t BigInteger) String() string {
	return "BIGINT"
}

type Text struct{}

func (t Text) Type() {}

func (t Text) String() string {
	return "TEXT"
}
