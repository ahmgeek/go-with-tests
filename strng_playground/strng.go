package strng

import "strings"

// Contains method; tests strings.Contains from "strings" package,
// just for fun.
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// ContainsAny method; tests strings.ContainsAny from "strings" package,
// Just for meh.
func ContainsAny(str, substr string) bool {
	return strings.ContainsAny(str, substr)
}
