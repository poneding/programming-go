package main

import (
	"fmt"
)

func longestPalindrome(str string) []string {
	l := len(str)
	if l == 0 || l == 1 {
		return []string{str}
	}

	res := make(map[int][]string)

	for i, _ := range str {
		oddLeft, oddRight, evenLeft, evenRight := i, i, i, i+1

		// 回文长度为奇数结果
		subStrLeft := helper(str, oddLeft, oddRight)

		// 回文长度为偶数结果
		subStrRight := helper(str, evenLeft, evenRight)

		if len(subStrLeft) > len(subStrRight) {
			res[len(subStrLeft)] = append(res[len(subStrLeft)], subStrLeft)
		} else {
			res[len(subStrRight)] = append(res[len(subStrRight)], subStrRight)
		}
	}

	max := 0
	for k, _ := range res {
		if k > max {
			max = k
		}
	}
	return res[max]
}

// 返回长度和回文串
func helper(str string, left, right int) string {
	subStr := str[left:right]
	for left >= 0 && right < len(str) && str[left] == str[right] {
		subStr = str[left : right+1]
		left--
		right++
	}
	return subStr
}

func main() {
	fmt.Println(longestPalindrome("babad"))             // 返回 bab、 aba
	fmt.Println(longestPalindrome("babbad"))            // 返回 abba
	fmt.Println(longestPalindrome("babadzzzzzzzzzzzc")) // 返回 zzzzzzzzzzz
}
