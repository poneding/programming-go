package main

import "fmt"

func main() {
	fmt.Println(isRotateString("abcd", "cdab"))
	fmt.Println(isRotateString("abcd", "cdba"))
}

func isRotateString(a, b string) bool {
	lena := len(a)
	lenb := len(b)
	if lena != lenb {
		return false
	}

	a2 := a + a
	for i := 0; i < lena; i++ {
		if a2[i:lena+i] == b {
			return true
		}
	}
	return false
}
