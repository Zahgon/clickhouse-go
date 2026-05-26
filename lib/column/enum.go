package column

func Enum(chType Type, name string) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

// to be updated below, when ranging over all index/enum values

const (
	enum8Type = "Enum8"
	E
	enum16Type = "Enum16"
)

func extractEnumNamedValues(chType Type) (typ string, values []string, indexes []int, valid bool) {
	_ = "STUB: not implemented"
	return "", nil, nil, false
}

// open bracket found, capture the type

// Ignore everything captured as non-enum type

// when inside a bracket, we can start capture value inside single quotes

// close the string and capture the value

// escape character, skip the next character

// capture optional index. `=` token is followed with an integer index

// find the end of the index, it's either a comma or a closing bracket

// capture the value and index when a comma or closing bracket is found

// if no index was found for current value, increment the value index
// e.g. Enum8('a','b') is equivalent to Enum8('a'=1,'b'=2)
// or Enum8('a'=3,'b') is equivalent to Enum8('a'=3,'b'=4)
// so if no index is provided, we increment the value index

// if the index is out of range, return

// Enum type must have at least one value
