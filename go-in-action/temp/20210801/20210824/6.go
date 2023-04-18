package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("dabbac"))
}

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 || l == 1 {
		return s
	}

	m := map[int][]string{}
	for i, _ := range s {
		oddLeft, oddRight, evenLeft, evenRight := i, i, i, i+1
		// 奇数
		p1 := palindromeString(s, oddLeft, oddRight)
		// 偶数
		p2 := palindromeString(s, evenLeft, evenRight)
		if len(p1) > len(p2) {
			m[len(p1)] = append(m[len(p1)], p1)
		} else {
			m[len(p2)] = append(m[len(p2)], p2)
		}
	}
	var max int
	for k, _ := range m {
		if k > max {
			max = k
		}
	}
	return m[max][0]
}

func palindromeString(s string, left, right int) string {
	res := s[left:right]
	l := len(s)
	for left >= 0 && right < l && s[left] == s[right] {
		res = s[left : right+1]
		left--
		right++
	}
	return res
}

func longestPalindrome2(s string) string {
	l := len(s)
	if l == 0 || l == 1 {
		return s
	}

	m := map[int][]string{}

	for i, _ := range s {
		oddLeft, oddRight, evenLeft, evenRight := i, i, i, i+1

		// 回文长度为奇数结果
		p1 := palindromeString2(s, oddLeft, oddRight)

		// 回文长度为偶数结果
		p2 := palindromeString2(s, evenLeft, evenRight)
		if len(p1) > len(p2) {
			m[len(p1)] = append(m[len(p1)], p1)
		} else {
			m[len(p2)] = append(m[len(p2)], p2)
		}
	}

	max := 0
	for k, _ := range m {
		if k > max {
			max = k
		}
	}
	return m[max][0]
}

// 返回 回文串
func palindromeString2(s string, left, right int) string {
	subStr := s[left:right]
	for left >= 0 && right < len(s) && s[left] == s[right] {
		subStr = s[left : right+1]
		left--
		right++
	}
	return subStr
}
