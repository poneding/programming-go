package main

// func main() {
// 	// >>1 除以2; >>2 除以2*2，也就是4
// 	// <<1 乘以2
// 	fmt.Println(13 << 2)
// 	fmt.Println(4 ^ 6)
// 	//arr := StrArr([]string{"pone", "poneding", "jay", "chou"})
// 	//sort.Sort(arr)
// 	//sort.Ints([]int{5, 4, 7, 3, 6})
// 	//fmt.Println(arr)
// }

type StrArr []string

func (arr StrArr) Len() int {
	return len(arr)
}

func (arr StrArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr StrArr) Less(i, j int) bool {
	return len(arr[i]) < len(arr[j])
}

//func (arr StrArr) Less(i, j int) bool {
//	return arr[i] < arr[j]
//}
