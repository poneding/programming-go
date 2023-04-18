package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}

/*
$ go run epoch.go
2023-02-12 23:15:58.589297 +0800 CST m=+0.000301376
1676214958
1676214958589
1676214958589297000
2023-02-12 23:15:58 +0800 CST
2023-02-12 23:15:58.589297 +0800 CST
*/
