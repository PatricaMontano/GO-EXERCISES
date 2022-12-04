package word

import (
	"strings"
)

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder

// UseCount return a map with word obtained from dividing a string
// by its spaces as the key, and the value is the number of times a word
// is repeated
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count return the count of word separate by spaces in string
func Count(s string) int {
	a := strings.Fields(s)
	return len(a)
	// write the code for this func
}
