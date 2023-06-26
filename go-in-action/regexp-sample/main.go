package main

import (
	"fmt"
	"regexp"
)

var (
	reg = regexp.MustCompile(`(\d+)ms`)
)

func main() {
	log := "api1 elapsed: 54ms, api2 elapsed: 23ms"
	res := reg.FindAllSubmatchIndex([]byte(log), -1)

	fmt.Println("res:", res)
	// 将打印 [[14 18 14 16] [34 38 34 36]]
	// 表示匹配到两组结果，每组有4个元素，
	// 前两个元素是匹配到的字符串的起始和结束位置，
	// 后两个元素是匹配到的子表达式的起始和结束位置

	for _, r := range res {
		fmt.Println(string(log[r[2]:r[3]]))
	}
}
