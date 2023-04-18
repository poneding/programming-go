package dog_fish_cat

import (
	"fmt"
)

func PrintDog(ch chan string) {
	go func() {
		select {
		case s := <-ch:
			if s == "dog" {
				fmt.Println("dog")
				ch <- "fish"
			}
		}
	}()
}
func PrintFish(ch chan string) {
	go func() {
		select {
		case s := <-ch:
			if s == "fish" {
				fmt.Println("fish")
				ch <- "cat"
			}
		}
	}()
}
func PrintCat(ch chan string) {
	go func() {
		select {
		case s := <-ch:
			if s == "cat" {
				fmt.Println("cat")
				ch <- "dog"
			}
		}
	}()
}

func Do() {
	ch := make(chan string)
	ch <- "dog"

	for i := 0; i < 100; i++ {
		println()
	}
}
