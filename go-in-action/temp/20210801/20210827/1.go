package main

import "fmt"

func main1() {
	fmt.Println(isDeformation("123", "132"))
	fmt.Println(isDeformation("1234", "132"))
	fmt.Println(isDeformation("adsf", "afsd"))
	fmt.Println(isDeformation("adsf", ""))
}

func isDeformation(a, b string) bool {
	lena := len(a)
	lenb := len(b)
	if lena != lenb || lena == 0 || lenb == 0 {
		return false
	}

	var m [256]int
	for _, i := range a {
		m[i]++
	}

	for _, j := range b {
		if m[j] == 0 {
			return false
		}
		m[j]--
	}
	return true
}
