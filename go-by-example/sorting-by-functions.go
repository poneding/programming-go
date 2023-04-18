package main

import (
	"fmt"
	"sort"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	// 按字符串长度大小排序
	return len(s[i]) < len(s[j])
}

/* $ go run sorting-by-functions.go
[kiwi peach banana]
*/
