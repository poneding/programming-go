package main

import "fmt"

func main() {
	e := Run()
	if e != nil {
		fmt.Println(e.Error())
	}
	if myerr, ok := e.(*MyError); ok {
		fmt.Println(myerr.Info)
	}
}

type MyError struct {
	Info string
}

func NewMyError(info string) *MyError {
	return &MyError{Info: info}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: %s\n", e.Info)
}

func Run() error {
	return NewMyError("just a myerror.")
}
