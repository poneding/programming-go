package main

import "fmt"

func main2() {
	arr := []int{1, 2}
	swap(arr, 0, 1)
	fmt.Println(arr[0], arr[1])
	a, b := new(int), new(int)
	*a = 10
	*b = 20
	swap1(*a, *b)
	fmt.Println(a, b)

}

func swap1(a, b int) {
	// 需要满足a和b不是同一个变量，否则会清吸为0
	//因为 n^n=0
	a = a ^ b
	b = a ^ b
	a = a ^ b
}

func swap(arr []int, a, b int) {
	// 需要满足a和b不是同一个变量，否则会清吸为0
	//因为 n^n=0
	arr[a] = arr[a] ^ arr[b]
	arr[b] = arr[a] ^ arr[b]
	arr[a] = arr[a] ^ arr[b]
}
