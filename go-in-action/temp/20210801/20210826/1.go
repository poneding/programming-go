package main

import "fmt"

// 最长公共子串
// 例如 13479，123474，则最长公共子串是347

func main1() {
	fmt.Println(longestCommonSubString("13479", "123474"))
	fmt.Println(longestCommonSubString("12345", "12345"))
	fmt.Println(longestCommonSubString("3", "123474"))
	fmt.Println(longestCommonSubString("13479", "1"))
	fmt.Println(longestCommonSubString("HelloWorld", "MyWord"))
}

func longestCommonSubString(a, b string) string {
	lenA := len(a)
	lenB := len(b)

	if lenA == 0 || lenB == 0 {
		return ""
	}

	dp := make([][]int, lenA)
	for i := 0; i < lenA; i++ {
		dp[i] = make([]int, lenB)
	}
	var (
		maxLen, maxI int
	)

	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			if a[i] == b[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			} else {
				dp[i][j] = 0
			}
			maxLen = max(maxLen, dp[i][j])

			if maxLen == dp[i][j] {
				maxI = i
			}
		}
	}

	return a[maxI-maxLen+1 : maxI+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
