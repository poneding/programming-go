package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello, 周杰伦"
	fmt.Println(len(str))                    //16 获取到的是字节长度，一个汉字是三个字节
	fmt.Println(utf8.RuneCountInString(str)) // 10
	fmt.Println(len([]rune(str)))            // 10
	fmt.Println(strings.Count(str, "")) // 11 需要减1，才是正确字符串的长度
	fmt.Println(bytes.Count([]byte(str), nil)) // 11
}
