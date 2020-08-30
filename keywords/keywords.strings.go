// Code generated by "stringer -type=Word -output=keywords.strings.go"; DO NOT EDIT.

package keywords

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ACCESS-1]
	_ = x[ACTION-2]
	_ = x[AGGREGATE-3]
	_ = x[AS-4]
	_ = x[BIGINT-5]
	_ = x[BOOLEAN-6]
	_ = x[BY-7]
	_ = x[BYTEA-8]
	_ = x[CASCADE-9]
	_ = x[CAST-10]
	_ = x[CHAR-11]
	_ = x[CHARACTER-12]
	_ = x[CHECK-13]
	_ = x[COLLATE-14]
	_ = x[COLLATION-15]
	_ = x[CONSTRAINT-16]
	_ = x[CONVERSION-17]
	_ = x[CREATE-18]
	_ = x[CROSS-19]
	_ = x[DATA-20]
	_ = x[DATABASE-21]
	_ = x[DATE-22]
	_ = x[DEC-23]
	_ = x[DECIMAL-24]
	_ = x[DEFAULT-25]
	_ = x[DEFERRABLE-26]
	_ = x[DEFERRED-27]
	_ = x[DELETE-28]
	_ = x[DISTINCT-29]
	_ = x[DOMAIN-30]
	_ = x[DOUBLE-31]
	_ = x[EVENT-32]
	_ = x[EXISTS-33]
	_ = x[EXTENSION-34]
	_ = x[FLOAT-35]
	_ = x[FOREIGN-36]
	_ = x[FROM-37]
	_ = x[FUNCTION-38]
	_ = x[GLOBAL-39]
	_ = x[GROUP-40]
	_ = x[HAVING-41]
	_ = x[IF-42]
	_ = x[IMMEDIATE-43]
	_ = x[INDEX-44]
	_ = x[INHERITS-45]
	_ = x[INITIALLY-46]
	_ = x[INNER-47]
	_ = x[INSERT-48]
	_ = x[INT-49]
	_ = x[INTEGER-50]
	_ = x[INTERVAL-51]
	_ = x[INTO-52]
	_ = x[IS-53]
	_ = x[JOIN-54]
	_ = x[KEY-55]
	_ = x[LEFT-56]
	_ = x[LIMIT-57]
	_ = x[LOCAL-58]
	_ = x[METHOD-59]
	_ = x[NO-60]
	_ = x[NOT-61]
	_ = x[NULL-62]
	_ = x[NUMERIC-63]
	_ = x[OFFSET-64]
	_ = x[ON-65]
	_ = x[OR-66]
	_ = x[OUTER-67]
	_ = x[PRIMARY-68]
	_ = x[REAL-69]
	_ = x[RECURSIVE-70]
	_ = x[REFERENCES-71]
	_ = x[REGCLASS-72]
	_ = x[REPLACE-73]
	_ = x[RESTRICT-74]
	_ = x[SCHEMA-75]
	_ = x[SELECT-76]
	_ = x[SET-77]
	_ = x[SMALLINT-78]
	_ = x[TABLE-79]
	_ = x[TEMP-80]
	_ = x[TEMPORARY-81]
	_ = x[TEXT-82]
	_ = x[TIME-83]
	_ = x[TIMESTAMP-84]
	_ = x[TRIGGER-85]
	_ = x[UNIQUE-86]
	_ = x[UNLOGGED-87]
	_ = x[UPDATE-88]
	_ = x[USER-89]
	_ = x[UUID-90]
	_ = x[VALUES-91]
	_ = x[VARCHAR-92]
	_ = x[VIEW-93]
	_ = x[WHERE-94]
	_ = x[WITH-95]
	_ = x[WRAPPER-96]
}

const _Word_name = "ACCESSACTIONAGGREGATEASBIGINTBOOLEANBYBYTEACASCADECASTCHARCHARACTERCHECKCOLLATECOLLATIONCONSTRAINTCONVERSIONCREATECROSSDATADATABASEDATEDECDECIMALDEFAULTDEFERRABLEDEFERREDDELETEDISTINCTDOMAINDOUBLEEVENTEXISTSEXTENSIONFLOATFOREIGNFROMFUNCTIONGLOBALGROUPHAVINGIFIMMEDIATEINDEXINHERITSINITIALLYINNERINSERTINTINTEGERINTERVALINTOISJOINKEYLEFTLIMITLOCALMETHODNONOTNULLNUMERICOFFSETONOROUTERPRIMARYREALRECURSIVEREFERENCESREGCLASSREPLACERESTRICTSCHEMASELECTSETSMALLINTTABLETEMPTEMPORARYTEXTTIMETIMESTAMPTRIGGERUNIQUEUNLOGGEDUPDATEUSERUUIDVALUESVARCHARVIEWWHEREWITHWRAPPER"

var _Word_index = [...]uint16{0, 6, 12, 21, 23, 29, 36, 38, 43, 50, 54, 58, 67, 72, 79, 88, 98, 108, 114, 119, 123, 131, 135, 138, 145, 152, 162, 170, 176, 184, 190, 196, 201, 207, 216, 221, 228, 232, 240, 246, 251, 257, 259, 268, 273, 281, 290, 295, 301, 304, 311, 319, 323, 325, 329, 332, 336, 341, 346, 352, 354, 357, 361, 368, 374, 376, 378, 383, 390, 394, 403, 413, 421, 428, 436, 442, 448, 451, 459, 464, 468, 477, 481, 485, 494, 501, 507, 515, 521, 525, 529, 535, 542, 546, 551, 555, 562}

func (i Word) String() string {
	i -= 1
	if i >= Word(len(_Word_index)-1) {
		return "Word(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Word_name[_Word_index[i]:_Word_index[i+1]]
}
