package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p（[a-z]）ch", "peach")
	fmt.Printf("match: %v\n", match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Printf("r.MatchString(\"peach\"): %v\n", r.MatchString("peach"))

	fmt.Printf("r.FindString(\"peach punch\"): %v\n", r.FindString("peach punch"))

	fmt.Printf("r.FindStringIndex(\"peach punch\"): %v\n", r.FindStringIndex("peach punch"))

	fmt.Printf("r.FindStringSubmatch(\"peach punch\"): %v\n", r.FindStringSubmatch("peach punch"))

	fmt.Printf("r.FindStringSubmatchIndex(\"peach punch\"): %v\n", r.FindStringSubmatchIndex("peach punch"))

	fmt.Printf("r.FindAllString(\"peach punch pinch\", -1): %v\n", r.FindAllString("peach punch pinch", -1))

	fmt.Printf("r.FindAllStringSubmatchIndex(\"peach punch pinch\", -1): %v\n", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Printf("r.FindAllString(\"peach punch pinch\", 2): %v\n", r.FindAllString("peach punch pinch", 2))

	fmt.Printf("r.Match([]byte(\"peach\")): %v\n", r.Match([]byte("peach")))

	fmt.Printf("r.ReplaceAllString(\"a peach\", \"<fruit>\"): %v\n", r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
