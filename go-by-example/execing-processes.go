package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	bin, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	err = syscall.Exec(bin, args, env)
	if err != nil {
		panic(err)
	}
}

/*
$ go run execing-processes.go
total 32
drwxr-xr-x  6 dp  staff   192B Jan 18 15:19 .
drwxr-xr-x  7 dp  staff   224B Jan 18 11:50 ..
-rw-r--r--  1 dp  staff   270B Jan 18 15:21 execing-processes.go
*/
