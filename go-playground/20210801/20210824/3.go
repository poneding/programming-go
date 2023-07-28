package main

import "fmt"

func main3() {
	fmt.Println(twoSum([]int{5, 2, 3, 2, 6}, 9))
	fmt.Println(twoSum([]int{2, 2}, 4))
}

//func twoSum1(nums []int, target int) []int {
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int, 0)

	for i, n := range nums {
		a := target - n

		if _, ok := m[a]; ok {
			res = []int{i, m[a]}
			break
		}
		m[n] = i
	}
	return res
}
