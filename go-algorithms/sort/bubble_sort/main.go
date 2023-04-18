package main

import "fmt"

func main() {
	nums := []int{5, 3, 7, 1, 9, 4, 5}
	bubbleSort(nums)

	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	l := len(nums)
	if l == 0 {
		return
	}

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				// 直接交换
				// nums[i], nums[j] = nums[j], nums[i]

				// 位运算交换
				// a ^ a = 0
				// 0 ^ n = n
				nums[i] = nums[i] ^ nums[j] // a^b
				nums[j] = nums[i] ^ nums[j] // (a^b)^b =>a
				nums[i] = nums[i] ^ nums[j] // a^(a^b) =>b
			}
		}
	}
}
