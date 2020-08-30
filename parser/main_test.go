package parser

import (
	"github.com/elliotcourant/pgoparser/tree"
	"github.com/stretchr/testify/require"
	"testing"
)

func MustParse(t *testing.T, query string) []tree.Statement {
	statements, err := Parse(query)
	require.NoErrorf(t, err, "%s - should have parse successfully", query)
	return statements
}
