package main

import (
	"errors"
	"fmt"
	"math"
)

//1. 实现字符串转整数
//实现java的parseInt函数，输入一个字符串返回一个整数。需要和parseInt行为一致。注意输入的字符不一定是整数，可能包含特殊字符或字母￼
func main() {
	fmt.Println(parseInt("123"))
	fmt.Println(parseInt("+123"))
	fmt.Println(parseInt("-123"))
	fmt.Println(parseInt("  -123.00"))
}

func parseInt(str string) (int, error) {
	var num int
	var flag bool // num is minus
	for i, r := range str {
		if i == 0 {
			if r == 45 { // first char: -
				flag = true
				continue
			} else if r == 43 { // first char: +
				continue
			}
		}
		if r >= 48 && r <= 57 {
			// int64 size limit
			if num >= math.MaxInt64/10 {
				mantissa := math.MaxInt64 % 10
				if flag {
					mantissa += 1
				}
				if int(r-48) > mantissa {
					return 0, errors.New("out of int64 size")
				}
			}
			num = num*10 + int(r-48)
		} else {
			return 0, errors.New("invalid string")
		}
	}
	fmt.Println("num here:", num)
	if flag {
		num = 0 - num
	}
	return num, nil
}
