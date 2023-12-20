package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 3, 6, 3, 9, 7, 3}
	selectionSort(nums)
	fmt.Println(nums)
}

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min == i {
			continue
		}
		// 直接交换两个数
		// nums[i], nums[min] = nums[min], nums[i]

		// 使用异或运算交换两个数
		// 注意： 两个数一定不能是同一个变量。
		// n^n=0	0^n=n
		nums[i] ^= nums[min] // a^b
		nums[min] ^= nums[i] // (a^b)^b =>a
		nums[i] ^= nums[min] // a^(a^b) =>b
	}
	return nums
}
