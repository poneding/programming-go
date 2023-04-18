package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "汉字"
	fmt.Printf("len(s): %v\n", len(s)) // 6

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\t", s[i]) // e6      b1      89      e5      ad      97
	}

	fmt.Println()

	for i, runeVal := range s {
		fmt.Printf("%#U starts at %d\n", runeVal, i)
	}

	fmt.Println("rune count:", utf8.RuneCountInString(s)) // 2

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == '汉' {
		fmt.Println("found 汉")
	}
}
