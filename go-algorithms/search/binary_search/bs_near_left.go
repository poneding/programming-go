package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 6, 8, 8, 8, 34, 51, 102}
	fmt.Println(bsNearLeft(nums, 8))
	fmt.Println(bsNearLeft(nums, 7))
	fmt.Println(bsNearLeft(nums, 35))
	fmt.Println(bsNearLeft(nums, 222))
	fmt.Println(bsNearLeft(nums, -1))
}

// 在arr上，找满足>=value的最左位置
func bsNearLeft(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right, mid, res := 0, len(nums)-1, 0, -1
	for right >= left {
		mid = left + (right-left)>>1
		if nums[mid] >= target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
