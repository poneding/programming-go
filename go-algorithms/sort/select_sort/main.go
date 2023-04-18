package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 3, 6, 1, 9, 7, 3}
	selectSort(nums)
	fmt.Println(nums)
}

func selectSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// 如果这里minIndex恰好是i，那么使用swap函数会有问题，所以需要作判断。
		// 其实最好都做判断，避免没必要的swap
		if i != minIndex {
			swap(nums, i, minIndex)
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
