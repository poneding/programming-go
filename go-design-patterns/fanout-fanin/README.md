# FanOut/FanIn 扇出扇入模式

使用扇出扇入模式可以将一个任务分解为多个子任务并行执行，然后将所有子任务的结果合并为一个结果，这种模式在并发编程中非常有用。

所以，扇出扇入模式有两个阶段：

- 扇出阶段：将一个任务分解为多个子任务并行执行。
- 扇入阶段：将所有子任务的结果合并为一个结果。

## 例子

*main.go*：

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	run(3)
}

func worker(id int, jobs <-chan int, result chan<- string) {
	for j := range jobs {
		fmt.Printf("Worker with id %d started job %d\n", id, j)
		time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		// fan-in
		result <- fmt.Sprintf("Worker with id %d finished job %d", id, j)
	}
}

func run(noOfWorkers int) {
    jobCount := 10
	jobs := make(chan int, jobCount)
	result := make(chan string, jobCount)

	for i := 0; i < noOfWorkers; i++ {
		go worker(i, jobs, result)
	}

	for j := 0; j < jobCount; j++ {
		// fan-out
		jobs <- j
	}
	close(jobs)

	for r := 0; r < jobCount; r++ {
		fmt.Println(<-result)
	}
}
```

## 运行

```bash
go run main.go
```
