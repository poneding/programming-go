package task

import (
	"fmt"
	"runtime"
	"testing"
)

type Alarm struct {
	Message string
}

func (a Alarm) Do() {
	fmt.Println("ding ding ding ~")
	fmt.Println(a.Message)
}

func TestTask(t *testing.T) {
	taskPool := NewTaskPool(100)
	taskPool.Open()

	go func() {
		for i := 0; i < 1000; i++ {
			taskPool.Accept(Alarm{
				Message: fmt.Sprintf("[%d] it's time to exercise.", i),
			})
			//time.Sleep(time.Second)
		}
	}()

	for {
		fmt.Printf("runtime.NumGoroutine: %d\n", runtime.NumGoroutine())
	}
}