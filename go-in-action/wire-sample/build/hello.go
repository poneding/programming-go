//go:build !greet
// +build !greet

package main

import "fmt"

func Greet(name string) {
	fmt.Println("Hello, " + name)
}
