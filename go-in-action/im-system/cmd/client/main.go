package main

import "goim"

func main() {
	goim.NewClient("127.0.0.1", 8081).Run()
}
