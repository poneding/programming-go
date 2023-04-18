package main

import (
	"fmt"
)

// 找出目标值在数组中的索引，假设一定存在
func main2() {
	nums := []int{1, 3, 7, 8, 12, 32, 34, 51, 102}
	fmt.Println(binarySearch(nums, 32))
	fmt.Println(binarySearch(nums, 21))
}

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return helper(nums, 0, len(nums)-1, target)
}

// 1 递归法
// func helper(nums []int, left, right, target int) int {
// 	if left > right {
// 		return -1 // 表示未找到目标值
// 	}

// 	mid := left + (right-left)>>1
// 	if target == nums[mid] {
// 		return mid
// 	} else if target > nums[mid] {
// 		return helper(nums, mid+1, right, target)
// 	} else {
// 		return helper(nums, left, mid-1, target)
// 	}
// }

// 2 迭代法
func helper(nums []int, left, right, target int) int {
	var mid int
	for right >= left {
		mid = left + (right-left)>>1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
