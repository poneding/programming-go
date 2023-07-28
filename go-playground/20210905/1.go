package main

import "fmt"

func main1() {
	// 异或运算
	// 1, 0^n=n	n^n=0
	// 2, a^b^c=b^c^a
	fmt.Println(twoOdd([]int{1, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 5}))
}

// 第一题：一个数组中只有一个数值出现了奇数次，其他数值都出现了偶数次，求奇数次值
func oneOdd(arr []int) int {
	var eor int
	for _, a := range arr {
		eor ^= a
	}
	return eor
}

//第一题：一个数组中只有2个数值出现了奇数次，其他数值都出现了偶数次，求奇数次值
func twoOdd(arr []int) (int, int) {
	var eor int
	for _, a := range arr {
		eor ^= a
	}

	var eor2 int
	right1 := rightOne(eor)
	for _, a := range arr {
		if a&right1 == 0 {
			eor2 ^= a
		}
	}
	return eor ^ eor2, eor2
}

// 一个二进制数提取最右侧的1
// 例如：
// 二进制值为：	10011000	a值
// 第一步取反：	01100111	a取反
// 加1			00000001	1值
// 得到值		01101000	a取反后加1
// 与a与运算		00001000	得到最右侧1
func rightOne(a int) int {
	return a & (^a + 1)
}
