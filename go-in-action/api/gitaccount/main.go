package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	//strings := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "s", "t", "r", "v", "w", "x", "y", "z"}
	strings := []string{"a", "i", "k", "l", "n", "o", "s", "r", "x", "y", "z"}
	var name string
	wg := &sync.WaitGroup{}

	for _, c1 := range strings {
		for _, c2 := range strings {
			name = "pone" + c1 + c2
			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				if IsGithubAccountSignUp(s) {
					fmt.Println("https://github.com/" + s)
				} else {
					fmt.Print("-")
				}
			}(name)
		}
	}

	wg.Wait()
}

func IsGithubAccountSignUp(name string) bool {
	resp, err := http.Get("https://github.com/" + name)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err.Error())
	}
	defer resp.Body.Close()
	res := resp.StatusCode == 404
	return res
}
