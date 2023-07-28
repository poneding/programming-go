package main

import "testing"

func FuzzOverwriteString(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string, val rune, n int) {
		OverwriteString(str, val, n)
	})
}
