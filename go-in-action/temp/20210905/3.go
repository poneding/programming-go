package main

import (
	"fmt"
	"math"
)

func main() {
	// 取中点得写法
	// 一般写法是 (r+l)/2，但是这种方法可能会int溢出，即（r+l）>int.max
	// 可以改写成 l +(r-l)/2
	// 除2可以使用位运算： >> 1
	//所以 可以使用 mid=l+(r-l)>>1
	r,l :=math.MaxInt64/2,math.MaxInt64
	//mid:=(r+l)  //存在溢出问题
	//mid:=l+(r-l)/2 // 可以
	mid:=l+(r-l)>>1
	fmt.Println(mid)
}
