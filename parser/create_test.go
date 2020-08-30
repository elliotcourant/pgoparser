package parser

import (
	"github.com/elliotcourant/pgoparser/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTable(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		result := MustParse(t, `CREATE TABLE users (id BIGINT PRIMARY KEY, email TEXT UNIQUE NOT NULL);`)
		assert.IsType(t, tree.CreateTableStatement{}, result[0])
	})

	t.Run("temp", func(t *testing.T) {
		result := MustParse(t, `CREATE TEMP TABLE users (id BIGINT PRIMARY KEY);`)
		assert.True(t, result[0].(tree.CreateTableStatement).Temporary, "temporary should be true")
	})

	t.Run("temporary", func(t *testing.T) {
		result := MustParse(t, `CREATE TEMPORARY TABLE users (id BIGINT PRIMARY KEY);`)
		assert.True(t, result[0].(tree.CreateTableStatement).Temporary, "temporary should be true")
	})

	t.Run("unlogged", func(t *testing.T) {
		result := MustParse(t, `CREATE UNLOGGED TABLE users (id BIGINT PRIMARY KEY);`)
		assert.True(t, result[0].(tree.CreateTableStatement).Unlogged, "unlogged should be true")
	})

	t.Run("if not exists", func(t *testing.T) {
		result := MustParse(t, `CREATE TABLE IF NOT EXISTS users (id BIGINT PRIMARY KEY);`)
		assert.True(t, result[0].(tree.CreateTableStatement).IfNotExists, "if not exists should be true")
	})

	t.Run("inherits", func(t *testing.T) {
		result := MustParse(t, `CREATE TABLE users (id BIGINT PRIMARY KEY) INHERITS accounts, billing;`)
		assert.Len(t, result[0].(tree.CreateTableStatement).Inherits, 2, "should have two inherited tables")
	})
}
