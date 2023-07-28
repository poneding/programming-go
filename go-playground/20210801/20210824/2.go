package main

import "fmt"

func main2() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("12512511035"))
	fmt.Println(restoreIpAddresses("11111"))
	fmt.Println(restoreIpAddresses("120010"))
}

func restoreIpAddresses(s string) []string {
	var res []string
	l := len(s)
	if l < 4 {
		return res
	}
	p1, p2, p3 := 1, 2, 3 // 三个切分点，代表.的位置
	for {
		// 切分之后的ip段
		s1, s2, s3, s4 := s[0:p1], s[p1:p2], s[p2:p3], s[p3:]
		if isValid(s1) && isValid(s2) && isValid(s3) && isValid(s4) {
			res = append(res, s1+"."+s2+"."+s3+"."+s4)
		}

		if p3 < l-1 { // p3往后移，直至无法再移动
			p3++
		} else if p2 < l-2 { // p2往后移，直至无法再移动，p3重置紧随p2之后
			p2, p3 = p2+1, p2+2
		} else if p1 < l-3 { // p1往后移，直至无法再移动，p2,p3重置紧随p1之后
			p1, p2, p3 = p1+1, p1+2, p1+3
		} else { //	p1,p2,p3都无法再移动，返回
			return res
		}
	}
}

// 判断一个字符是否IP段合法
// "0","1","255"合法
// "00","01","256"不合法
func isValid(s string) bool {
	if s[0] == '0' {
		return s == "0"
	}
	var num int
	for _, r := range s {
		num = num*10 + int(r) - 48
	}

	return num <= 255
}
