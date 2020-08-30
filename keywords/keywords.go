package keywords

import (
	"github.com/elliotcourant/pgoparser/quotes"
	"github.com/elliotcourant/pgoparser/tokens"
	"github.com/elliotcourant/pgoparser/words"
	"strings"
)

type Keyword interface {
	words.Word
	Keyword()
}

var (
	_ Keyword = Word(0)
)

//go:generate stringer -type=Word -output=keywords.strings.go
type Word uint8

func (i Word) Token() {}

func (i Word) Word() {}

func (i Word) Quotes() quotes.Quotes {
	return quotes.None
}

func (i Word) Keyword() {}

func IsKeyword(keyword string) bool {
	_, ok := _keywordMap[strings.ToUpper(keyword)]
	return ok
}

func NewKeywordMaybe(keyword string) tokens.Token {
	word, ok := _keywordMap[strings.ToUpper(keyword)]
	if !ok {
		return words.NewWord(keyword, quotes.None)
	}

	return word
}

const (
	_ Word = iota
	ACCESS
	ACTION
	AGGREGATE
	AS
	BIGINT
	BOOLEAN
	BY
	BYTEA
	CASCADE
	CAST
	CHAR
	CHARACTER
	CHECK
	COLLATE
	COLLATION
	CONSTRAINT
	CONVERSION
	CREATE
	CROSS
	DATA
	DATABASE
	DATE
	DEC
	DECIMAL
	DEFAULT
	DEFERRABLE
	DEFERRED
	DELETE
	DISTINCT
	DOMAIN
	DOUBLE
	EVENT
	EXISTS
	EXTENSION
	FLOAT
	FOREIGN
	FROM
	FUNCTION
	GLOBAL
	GROUP
	HAVING
	IF
	IMMEDIATE
	INDEX
	INHERITS
	INITIALLY
	INNER
	INSERT
	INT
	INTEGER
	INTERVAL
	INTO
	IS
	JOIN
	KEY
	LEFT
	LIMIT
	LOCAL
	METHOD
	NO
	NOT
	NULL
	NUMERIC
	OFFSET
	ON
	OR
	OUTER
	PRIMARY
	REAL
	RECURSIVE
	REFERENCES
	REGCLASS
	REPLACE
	RESTRICT
	SCHEMA
	SELECT
	SET
	SMALLINT
	TABLE
	TEMP
	TEMPORARY
	TEXT
	TIME
	TIMESTAMP
	TRIGGER
	UNIQUE
	UNLOGGED
	UPDATE
	USER
	UUID
	VALUES
	VARCHAR
	VIEW
	WHERE
	WITH
	WRAPPER
)
