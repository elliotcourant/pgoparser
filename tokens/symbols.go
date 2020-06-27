package tokens

type Comma struct{}

func (s Comma) Token() {}

func (s Comma) String() string { return "," }

type Period struct{}

func (s Period) Token() {}

func (s Period) String() string { return "." }

type SemiColon struct{}

func (s SemiColon) Token() {}

func (s SemiColon) String() string { return ";" }

type Equals struct{}

func (s Equals) Token() {}

func (s Equals) String() string { return "=" }

type NotEquals struct {
	Value string
}

func (s NotEquals) Token() {}

func (s NotEquals) String() string { return s.Value }

type LessThan struct{}

func (s LessThan) Token() {}

func (s LessThan) String() string { return "<" }

type GreaterThan struct{}

func (s GreaterThan) Token() {}

func (s GreaterThan) String() string { return ">" }

type LessThanOrEqualTo struct{}

func (s LessThanOrEqualTo) Token() {}

func (s LessThanOrEqualTo) String() string { return "<=" }

type GreaterThanOrEqualTo struct{}

func (s GreaterThanOrEqualTo) Token() {}

func (s GreaterThanOrEqualTo) String() string { return ">=" }

type LeftParentheses struct{}

func (s LeftParentheses) Token() {}

func (s LeftParentheses) String() string { return "(" }

type RightParentheses struct{}

func (s RightParentheses) Token() {}

func (s RightParentheses) String() string { return ")" }

type Plus struct{}

func (s Plus) Token() {}

func (s Plus) String() string { return "+" }

type Minus struct{}

func (s Minus) Token() {}

func (s Minus) String() string { return "-" }

type Division struct{}

func (s Division) Token() {}

func (s Division) String() string { return "/" }

type Multiply struct{}

func (s Multiply) Token() {}

func (s Multiply) String() string { return "*" }

type Modulo struct{}

func (s Modulo) Token() {}

func (s Modulo) String() string { return "%" }

type Pipe struct{}

func (s Pipe) Token() {}

func (s Pipe) String() string { return "|" }

type StringConcatenation struct{}

func (s StringConcatenation) Token() {}

func (s StringConcatenation) String() string { return "||" }
