package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	os.Exit(3)
}

/*
$ go run exit.go
exit status 3

$ go build exit.go
$ ./exit
$ echo $?
3
*/
