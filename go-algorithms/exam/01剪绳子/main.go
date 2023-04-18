package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(cuttingRope(10))
	fmt.Println(cuttingRope2(10))
}

// 数学推导而得的规律
func cuttingRope(n int) int {

	res := 1
	if n >= 4 {
		a := n / 3
		b := n % 3
		if b == 1 {
			a -= 1
			res *= 4
		} else if b == 2 {
			res *= 2
		}

		for i := 0; i < a; i++ {
			res *= 3
		}
	} else {
		return n - 1
	}
	return res
}

// 动态规划
func cuttingRope2(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] = int(math.Max(float64(dp[i-j]*j), float64((i-j)*j)))
			if dp[i] < i {
				dp[i] = i
			}
		}
	}
	return dp[n-1]
}
