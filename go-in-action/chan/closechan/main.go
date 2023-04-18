package main

import (
	"fmt"
	"time"
)

func main() {
	// main1()
	main2()
}

func count(msg string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- msg
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // 发送端需要关闭chan
}

func main1() {
	c := make(chan string)
	go func() {
		count("sheep", c)
	}()

	// for {
	// 	msg := <-c
	// 	fmt.Println(msg)
	// }

	// for {
	// 	msg, opened := <-c // 发送端关闭了chan后，这里才能正常的使用!opened
	// 	if !opened {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range c {
		fmt.Println(msg)
	}
}

func main2() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "every 500 milliseconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for {
		// 因为chan没有缓冲，所以下面会堵塞
		// fmt.Println(<-c1)
		// fmt.Println(<-c2)

		// 这种方式可以正常打印
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}
