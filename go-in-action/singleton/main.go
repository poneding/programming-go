package main

import (
	"fmt"
	"time"
)

func main() {
	i1_1 := Instance()
	fmt.Printf("i1_1.Id: %s\n", i1_1.Id)

	time.Sleep(time.Second)
	i1_2 := Instance()
	fmt.Printf("i1_2.Id: %s\n", i1_2.Id)

	i2_1 := Instance2()
	fmt.Printf("i2_1.Id: %s\n", i2_1.Id)

	time.Sleep(time.Second)
	i2_2 := Instance2()
	fmt.Printf("i2_2.Id: %s\n", i2_2.Id)
}

