package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func main() {
// 	run()
// }

func run() {
	// 个输出结果决定来⾃于调度器优先调度哪个G。从runtime的源码可以看到，当创建⼀
	//个G时，会优先放⼊到下⼀个调度的 runnext 字段上作为下⼀次优先调度的G。因此，
	//最先输出的是最后创建的G，也就是19.
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		go func() {
			fmt.Println("A:", i)
			wg.Done()
		}()
	}

	for i := 10; i < 20; i++ {
		go func(i int) {
			fmt.Println("B:", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
