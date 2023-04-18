package main

import (
	"fmt"
	"strconv"
)

var letterMap = map[int][]string{
	1: {"a", "b"},
	2: {"c", "f", "e"},
	3: {"f", "g"},
}

func main() {
	res := CombineLetters("12")
	fmt.Println(res)
	fmt.Println(len(res))
}

// 传入已经组合好的
func combineLetters(nums string, prefix []string) []string {
	if len(nums) == 0 {
		return prefix
	}
	mapKey, _ := strconv.Atoi(nums[:1])
	if len(prefix) == 0 {
		prefix = letterMap[mapKey]
	} else {
		prefixLen := len(prefix)
		for _, p := range prefix {
			for _, l := range letterMap[mapKey] {
				prefix = append(prefix, p+l)
			}
		}
		prefix = prefix[prefixLen:]
	}

	res := combineLetters(nums[1:], prefix)
	return res
}

func CombineLetters(nums string) []string {
	res := combineLetters(nums, []string{})
	return res
}
