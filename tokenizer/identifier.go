package tokenizer

// isIdentifierStart will return true if the provided character is one of the allowed characters in the start of an
// identifier.
func (t *Tokenizer) isIdentifierStart(character byte) bool {
	// See https://www.postgresql.org/docs/13/sql-syntax-lexical.html
	// and https://www.postgresql.org/docs/11/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
	// We don't yet support identifiers beginning with "letters with diacritical marks and non-Latin letters"
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') || character == '_'
}

// isIdentifierPart will return true if the provided character is allowed to make up part of a PostgreSQL identifer.
func (t *Tokenizer) isIdentifierPart(character byte) bool {
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') ||
		(character >= '0' && character <= '9') ||
		(character == '$') ||
		(character == '_')
}
