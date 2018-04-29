package strng

import "strings"

// Contain method; tests strings.Contains from "strings" package,
// just for fun.
func Contain(word, target string) bool {
	return strings.Contains(word, target)
}
