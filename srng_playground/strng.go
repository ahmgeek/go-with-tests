package strng

import "strings"

func Contain(word, target string) bool {
	return strings.Contains(word, target)
}
