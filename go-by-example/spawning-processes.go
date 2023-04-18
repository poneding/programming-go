package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// date
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// date -x
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exec rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// grep
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()

	grepIn.Write([]byte("hello grep\nworld grep"))
	grepIn.Close()

	grepBytes, _ := io.ReadAll(grepOut)

	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// ls -a -l -h
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

/*
go run spawning-processes.go
> date
Wed Jan 18 15:41:37 CST 2023

command exec rc = 1
> grep hello
hello grep

> ls -a -l -h
total 40
drwxr-xr-x  7 dp  staff   224B Jan 18 15:24 .
drwxr-xr-x  7 dp  staff   224B Jan 18 11:50 ..
-rw-r--r--  1 dp  staff   1.0K Jan 18 15:41 spawning-processes.go
*/
