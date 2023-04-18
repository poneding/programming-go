package main

import (
	"strings"
)

func main() {
}

// overwrite_string.go

// OverwriteString overwrites the first 'n' characters in a string with
// the rune 'value'
func OverwriteString(str string, value rune, n int) string {
	// If asked to overwrite more than the entire string then no need to loop,
	// just return string length * the rune
	if n > len(str) {
		return strings.Repeat(string(value), len(str))
	}

	result := []rune(str)
	for i := 0; i <= n; i++ {
		result[i] = value
	}
	return string(result)
}
