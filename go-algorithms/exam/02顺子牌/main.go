package main

import (
	"fmt"
	"sort"
)

func main() {
	isStraight([]int{0, 0, 1, 2, 5})
}

func isStraight(nums []int) bool {
	var joker int
	sort.Ints(nums)
	fmt.Println(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			joker++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}

	}
	return nums[len(nums)-1]-nums[joker] < 5
}
