package main

import "fmt"

// 最长公共子序列

func main2() {
	fmt.Println(lcs("BDCABA", "ABCBDAB"))
	fmt.Println(lcs("", "ABCBDAB"))
	fmt.Println(lcs("BDCABA", ""))
	fmt.Println(lcs("ABCDE", "ABCDE"))
}

func lcs(a, b string) int {
	lena := len(a)
	lenb := len(b)
	if lena == 0 || lenb == 0 {
		return 0
	}

	dp := make([][]int, lena)
	for i := 0; i < lena; i++ {
		dp[i] = make([]int, lenb)
	}

	for i := 0; i < lena; i++ {
		for j := 0; j < lenb; j++ {
			if a[i] == b[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			} else {
				if i == 0 || j == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = max2(dp[i-1][j], dp[i][j-1])
				}
			}
		}
	}

	return dp[lena-1][lenb-1]
}

func max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}
