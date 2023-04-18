package main

import "fmt"

// abcabcbb

func main5() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("adcdabcdefa"))
}

func lengthOfLongestSubstring(s string) int {
	lens := len(s)
	if lens == 1 || len(s) == 0 {
		return lens
	}

	var (
		right int
		res   int
		m     = map[byte]int{}
	)

	for i := 0; i < lens; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for ; right < lens && m[s[right]] == 0; right++ { //上一轮符合条件时，这里已经+1
			m[s[right]]++
		}
		res = max(res, right-i) // 所以这里不能+1
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//func lengthOfLongestSubstring2(s string) int {
//	// 哈希集合，记录每个字符是否出现过
//	m := map[byte]int{}
//	n := len(s)
//	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
//	rk, ans := -1, 0
//	for i := 0; i < n; i++ {
//		if i != 0 {
//			// 左指针向右移动一格，移除一个字符
//			delete(m, s[i-1])
//		}
//		for rk + 1 < n && m[s[rk+1]] == 0 {
//			// 不断地移动右指针
//			m[s[rk+1]]++
//			rk++
//		}
//		// 第 i 到 rk 个字符是一个极长的无重复字符子串
//		ans = max2(ans, rk - i + 1)
//	}
//	return ans
//}
//
//func max2(x, y int) int {
//	if x < y {
//		return y
//	}
//	return x
//}
//
//func lengthOfLongestSubstring(s string) int {
//	m := make(map[byte]int, 0)
//	l := len(s)
//
//	right, ans := -1, 0
//	for i := 0; i < l; i++ {
//		if i != 0 {
//			delete(m, s[i-1])
//		}
//
//		for right+1 < l && m[s[right+1]] == 0 {
//			m[s[right+1]]++
//			right++
//		}
//
//		ans = max(ans, right-i+1)
//	}
//	return ans
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
