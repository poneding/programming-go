package main

import (
	"fmt"
	"strings"
)

// 行转列
// 示例：
// 假设file.txt文件内容如下：
// name age id
// alice 21 1
// ryan 30 2
//
// 应当输出：
// id 1 2
// name alice ryan
// age 21 30
func main() {
	s := `name age id
alice 21 1
ryan 30 2`
	lines := strings.Split(s, "\r\n")

	var array = make([][]string, len(lines))
	for i, line := range lines {
		array[i] = strings.Split(line, " ")
	}

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Print(array[j][i] + " ")
		}
		fmt.Println()
	}
}
