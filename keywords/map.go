package keywords

var (
	_keywordMap = map[string]Keyword{}
)

func init() {
	for i := 0; i < len(_Word_index)-1; i++ {
		str := _Word_name[_Word_index[i]:_Word_index[i+1]]
		_keywordMap[str] = Word(i + 1)
	}
}
