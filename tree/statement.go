package tree

type Statement interface {
	Statement()

	String() string
}
