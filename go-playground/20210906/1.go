package main

import "fmt"

// 求数组的最大子序和数组
// 例如： [-1,2,-3,4,-1,2,-3,2]，则最大子序和为[4,-1,2]
func main() {
	maxSum:=maxSubArraySum([]int{-1,2,-3,4,-1,2,-3,2})
	fmt.Println(maxSum)
}

// 解题思路1：暴力法
func maxSubArraySum(arr []int) int {
	var subArrSums []int
	for i := 0; i < len(arr); i++ {
		for j := i; j <len(arr); j++ {
			var  sum int
			for k := j; k < len(arr); k++ {
				sum += arr[j]
			}
			subArrSums=append(subArrSums,sum)
		}
	}

	var maxSum int
	for _,sum:=range subArrSums{
		maxSum=max(sum,maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
