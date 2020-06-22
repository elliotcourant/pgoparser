package keywords

import (
	"github.com/elliotcourant/pgoparser/words"
	"strings"
)

//go:generate keywordgen --input=./keywords.yaml --keyword-template=./templates/keywords.template --keyword-map-template=./templates/keyword_map.template
type Keyword interface {
	words.Word
	Keyword()
}

func NewKeyword(value string) Keyword {
	newKeyword, ok := keywordMap[strings.ToUpper(value)]
	if !ok {
		panic(value + " is not a valid keyword")
	}

	return newKeyword(value)
}

func IsValidKeyword(value string) bool {
	_, ok := keywordMap[strings.ToUpper(value)]
	return ok
}
