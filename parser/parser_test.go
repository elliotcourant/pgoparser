package parser

import (
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		sql := `CREATE TABLE users (id BIGINT PRIMARY KEY, email TEXT);`
		parsed, err := Parse(sql)
		assert.NoError(t, err)
		assert.NotNil(t, parsed)
	})
}

func TestStructMatch(t *testing.T) {
	var someToken tokens.Token
	someToken = tokens.EOF{}
	assert.True(t, someToken == (tokens.EOF{}))
	assert.False(t, someToken == (tokens.Comma{}))
}
