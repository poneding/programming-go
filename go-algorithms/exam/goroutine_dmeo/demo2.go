package main

import (
	"fmt"
	"log"
	"sync"
)

type LetterFreq map[rune]int

func countLetters(strs []string) LetterFreq {
	wg := sync.WaitGroup{}
	m := make(map[rune]int)
	ch := make(chan rune)

	for _, str := range strs {
		wg.Add(1)
		go helper(str, ch, &wg)
	}

	go func() {
		for {
			select {
			case r := <-ch:
				m[r]++
			}
		}
	}()

	wg.Wait()

	return LetterFreq(m)
}
func helper(str string, ch chan rune, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, r := range str {
		ch <- r
	}
}

func main() {
	inp := []string{"abc", "bcd", "aaa"}
	res := countLetters(inp)
	fmt.Println(res)

	if res['a'] == 4 {
		log.Println("OK")
	} else {
		log.Fatalf("No")
	}
}
