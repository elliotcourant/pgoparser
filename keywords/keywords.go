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
	SELECT Word = iota
	FROM
	WHERE
	LIMIT
	OFFSET
	CREATE
	TABLE
	PRIMARY
	KEY
	UNIQUE
	REFERENCES
	INDEX
	CONSTRAINT
	FOREIGN
	NULL
	DEFAULT
	IS
	NOT
	DISTINCT
	GROUP
	BY
	WITH
	HAVING
	INNER
	LEFT
	OUTER
	JOIN
	CROSS
	VIEW
	DATABASE
	USER
	AS
	SCHEMA
	IF
	EXISTS
	CHECK
	COLLATE
	BOOLEAN
	FLOAT
	REAL
	DOUBLE
	SMALLINT
	INT
	INTEGER
	BIGINT
	VARCHAR
	CHAR
	CHARACTER
	UUID
	DATE
	TIMESTAMP
	TIME
	INTERVAL
	REGCLASS
	TEXT
	BYTEA
	NUMERIC
	DECIMAL
	DEC
	ON
	DELETE
	UPDATE
	RESTRICT
	CASCADE
	SET
	ACTION
	NO
)
