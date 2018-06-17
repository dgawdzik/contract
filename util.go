package contract

import "strings"

/*
	Package Constants
*/

const (
	emptyString = ""
)

/*
	Public Functions
*/

// IsEmpty returns true when string is empty i.e. "" and false otherwise.
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == emptyString
}
