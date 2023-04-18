package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 3, 6, 1, 9, 7, 3}
	insertSort(nums)
	fmt.Println(nums)
}

func insertSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	// 先把前面的排号顺序，再依次将后面排进来
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0 && nums[j] > nums[j+1]; j-- {
			swap(nums, j, j+1)
		}
	}
}

// 注意： 两个数一定不能是同一个变量。
// n^n=0	0^n=n
func swap(nums []int, i, j int) {
	nums[i] ^= nums[j] // a^b
	nums[j] ^= nums[i] // (a^b)^b =>a
	nums[i] ^= nums[j] // a^(a^b) =>b
}
